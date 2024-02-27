package gin

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/LucasMateus-eng/operations-service/internal/logging"
	"github.com/LucasMateus-eng/operations-service/user"
	"github.com/gin-gonic/gin"
)

type userOutputDTO struct {
	ID             int       `json:"id"`
	Username       string    `json:"username,omitempty"`
	HashedPassword string    `json:"hashed_password,omitempty"`
	Role           user.Role `json:"role,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	DeletedAt      time.Time `json:"deleted_at,omitempty"`
}

type userInputDTO struct {
	ID             int       `json:"id"`
	Username       string    `json:"username" binding:"required"`
	HashedPassword string    `json:"hashed_password" binding:"required"`
	Role           user.Role `json:"role" binding:"required"`
}

func mapUserToOutputDTO(user user.User) *userOutputDTO {
	return &userOutputDTO{
		ID:             user.ID,
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		Role:           user.Role,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		DeletedAt:      user.DeletedAt,
	}
}

func mapInputDTOToUser(input userInputDTO) *user.User {
	return &user.User{
		ID:             input.ID,
		Username:       input.Username,
		HashedPassword: input.HashedPassword,
		Role:           input.Role,
	}
}

func getUser(ctx context.Context, service *user.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := service.GetById(ctx, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		outputDTO := mapUserToOutputDTO(*user)

		c.JSON(http.StatusOK, outputDTO)
	}
}

func createUser(ctx context.Context, service *user.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto userInputDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := mapInputDTOToUser(dto)

		userID, err := service.Create(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, userOutputDTO{ID: userID})
	}
}

func updateUser(ctx context.Context, service *user.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dto userInputDTO
		if err = c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := mapInputDTOToUser(dto)
		user.ID = userID

		err = service.Update(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func deleteUser(ctx context.Context, service *user.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = service.Delete(ctx, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
