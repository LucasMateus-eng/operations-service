package gin

import (
	"context"
	"net/http"
	"strconv"

	gin_dto "github.com/LucasMateus-eng/operations-service/internal/http/gin/dto"
	gin_mapping "github.com/LucasMateus-eng/operations-service/internal/http/gin/mapping"
	"github.com/LucasMateus-eng/operations-service/internal/logging"
	"github.com/LucasMateus-eng/operations-service/user"
	"github.com/gin-gonic/gin"
)

func getUser(ctx context.Context, service *user.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Get user", nil)

		userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := service.GetById(ctx, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		outputDTO := gin_mapping.MapUserToOutputDTO(*user)

		c.JSON(http.StatusOK, outputDTO)
	}
}

func createUser(ctx context.Context, service *user.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Create user", nil)

		var dto gin_dto.UserInputDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := gin_mapping.MapInputDTOToUser(dto)

		userID, err := service.Create(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin_dto.UserOutputDTO{ID: userID})
	}
}

func updateUser(ctx context.Context, service *user.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Update user", nil)

		userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dto gin_dto.UserInputDTO
		if err = c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := gin_mapping.MapInputDTOToUser(dto)
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
		logger.Info("Delete user", nil)

		userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
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
