package mapping

import (
	"testing"
	"time"

	"github.com/LucasMateus-eng/operations-service/address"
	address_dto "github.com/LucasMateus-eng/operations-service/address/postgres/dto"
	"github.com/go-playground/assert/v2"
)

var (
	mockedTime = time.Now()
)

func TestMapAddressToDTO(t *testing.T) {
	address := &address.Address{
		ID:           1,
		UserID:       1,
		Locality:     "Rua Principal",
		Number:       "123",
		Complement:   "Apartamento 101",
		Neighborhood: "Centro",
		City:         "Cidade",
		State:        address.BrazilianState(24),
		CEP:          "12345-678",
		Country:      "Brasil",
		CreatedAt:    mockedTime,
		UpdatedAt:    mockedTime,
		DeletedAt:    mockedTime,
	}

	expectedDTO := &address_dto.AddressDTO{
		ID:           1,
		Locality:     "Rua Principal",
		Number:       "123",
		Complement:   "Apartamento 101",
		Neighborhood: "Centro",
		City:         "Cidade",
		State:        "SÃO PAULO",
		CEP:          "12345-678",
		Country:      "Brasil",
		UserID:       1,
		CreatedAt:    address.CreatedAt,
		UpdatedAt:    address.UpdatedAt,
		DeletedAt:    address.DeletedAt,
	}

	actualDTO := MapAddressToDTO(address)
	assert.Equal(t, expectedDTO, actualDTO)
}

func TestMapDTOToAddress(t *testing.T) {
	addressDTO := &address_dto.AddressDTO{
		ID:           1,
		Locality:     "Rua Principal",
		Number:       "123",
		Complement:   "Apartamento 101",
		Neighborhood: "Centro",
		City:         "Cidade",
		State:        "SÃO PAULO",
		CEP:          "12345-678",
		Country:      "Brasil",
		UserID:       1,
		CreatedAt:    mockedTime,
		UpdatedAt:    mockedTime,
		DeletedAt:    mockedTime,
	}

	addressDTOWithInvalidState := &address_dto.AddressDTO{
		ID:           1,
		Locality:     "Rua Principal",
		Number:       "123",
		Complement:   "Apartamento 101",
		Neighborhood: "Centro",
		City:         "Cidade",
		State:        "MISSOURI",
		CEP:          "12345-678",
		Country:      "Brasil",
		UserID:       1,
		CreatedAt:    mockedTime,
		UpdatedAt:    mockedTime,
		DeletedAt:    mockedTime,
	}

	expectedAddress := &address.Address{
		ID:           1,
		Locality:     "Rua Principal",
		Number:       "123",
		Complement:   "Apartamento 101",
		Neighborhood: "Centro",
		City:         "Cidade",
		State:        address.BrazilianState(24),
		CEP:          "12345-678",
		Country:      "Brasil",
		UserID:       1,
		CreatedAt:    addressDTO.CreatedAt,
		UpdatedAt:    addressDTO.UpdatedAt,
		DeletedAt:    addressDTO.DeletedAt,
	}

	tests := []struct {
		name    string
		arg     *address_dto.AddressDTO
		want    *address.Address
		wantErr bool
	}{
		{
			name:    "Dado um DTO de Address quando a função de mapeamento é chamada então a conversão é um sucesso",
			arg:     addressDTO,
			want:    expectedAddress,
			wantErr: false,
		},
		{
			name:    "Dado um DTO de Address com um estado inválido quando a função de mapeamento é chamada então a conversão falha",
			arg:     addressDTOWithInvalidState,
			want:    nil,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			actualAddress, err := MapDTOToAddress(test.arg)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualAddress)
		})
	}
}
