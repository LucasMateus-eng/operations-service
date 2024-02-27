package gin

import (
	"time"

	"github.com/LucasMateus-eng/operations-service/user"
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

type userInputputDTO struct {
	ID             int       `json:"id"`
	Username       string    `json:"username" binding:"required"`
	HashedPassword string    `json:"hashed_password" binding:"required"`
	Role           user.Role `json:"role" binding:"required"`
}

func mapUserToOutputDTO(user *user.User) *userOutputDTO {
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

func mapInputDTOToUser(input *userInputputDTO) *user.User {
	return &user.User{
		ID:             input.ID,
		Username:       input.Username,
		HashedPassword: input.HashedPassword,
		Role:           input.Role,
	}
}
