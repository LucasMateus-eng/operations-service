package mapping

import (
	address_dto "github.com/LucasMateus-eng/operations-service/address/postgres/dto"
	"github.com/LucasMateus-eng/operations-service/user"
	user_dto "github.com/LucasMateus-eng/operations-service/user/postgres/dto"
)

func MapUserToDTO(user *user.User) *user_dto.UserDTO {
	return &user_dto.UserDTO{
		ID:             user.ID,
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		Role:           user.Role.String(),
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		DeletedAt:      user.DeletedAt,
		AddressDTO:     &address_dto.AddressDTO{},
	}
}

func MapDTOToUser(userDTO *user_dto.UserDTO) (*user.User, error) {
	role, err := user.GetRole(userDTO.Role)
	if err != nil {
		return nil, err
	}

	return &user.User{
		ID:             userDTO.ID,
		Username:       userDTO.Username,
		HashedPassword: userDTO.HashedPassword,
		Role:           *role,
		CreatedAt:      userDTO.CreatedAt,
		UpdatedAt:      userDTO.UpdatedAt,
		DeletedAt:      userDTO.DeletedAt,
	}, nil
}
