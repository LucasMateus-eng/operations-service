package mapping

import (
	"testing"
	"time"

	address_dto "github.com/LucasMateus-eng/operations-service/address/postgres/dto"
	"github.com/LucasMateus-eng/operations-service/user"
	user_dto "github.com/LucasMateus-eng/operations-service/user/postgres/dto"
	"github.com/go-playground/assert/v2"
)

var (
	mockedTime = time.Now()
)

func TestMapUserToDTO(t *testing.T) {
	user := &user.User{
		ID:             1,
		Username:       "user123",
		HashedPassword: "hashedpassword123",
		Role:           user.Role(1),
		CreatedAt:      mockedTime,
		UpdatedAt:      mockedTime,
		DeletedAt:      mockedTime,
	}

	expectedDTO := &user_dto.UserDTO{
		ID:             1,
		Username:       "user123",
		HashedPassword: "hashedpassword123",
		Role:           "ADMINISTRATOR",
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		DeletedAt:      user.DeletedAt,
		AddressDTO:     &address_dto.AddressDTO{},
	}

	actualDTO := MapUserToDTO(user)
	assert.Equal(t, expectedDTO, actualDTO)
}

func TestMapDTOToUser(t *testing.T) {
	userDTO := &user_dto.UserDTO{
		ID:             1,
		Username:       "user123",
		HashedPassword: "hashedpassword123",
		Role:           "ADMINISTRATOR",
		CreatedAt:      mockedTime,
		UpdatedAt:      mockedTime,
		DeletedAt:      mockedTime,
		AddressDTO:     &address_dto.AddressDTO{},
	}

	userDTOWithInvalidRole := &user_dto.UserDTO{
		ID:             1,
		Username:       "user123",
		HashedPassword: "hashedpassword123",
		Role:           "ENGINEER",
		CreatedAt:      mockedTime,
		UpdatedAt:      mockedTime,
		DeletedAt:      mockedTime,
		AddressDTO:     &address_dto.AddressDTO{},
	}

	expectedUser := &user.User{
		ID:             1,
		Username:       "user123",
		HashedPassword: "hashedpassword123",
		Role:           user.Role(1),
		CreatedAt:      userDTO.CreatedAt,
		UpdatedAt:      userDTO.UpdatedAt,
		DeletedAt:      userDTO.DeletedAt,
	}

	tests := []struct {
		name    string
		arg     *user_dto.UserDTO
		want    *user.User
		wantErr bool
	}{
		{
			name:    "Dado um DTO de User quando a função de mapeamento é chamada então a conversão é um sucesso",
			arg:     userDTO,
			want:    expectedUser,
			wantErr: false,
		},
		{
			name:    "Dado um DTO de User com uma função inválida quando a função de mapeamento é chamada então a conversão falha",
			arg:     userDTOWithInvalidRole,
			want:    nil,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			actualUser, err := MapDTOToUser(test.arg)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualUser)
		})
	}
}
