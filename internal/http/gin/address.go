package gin

import (
	"time"

	"github.com/LucasMateus-eng/operations-service/address"
)

type addressOutputDTO struct {
	ID           int                    `json:"id"`
	UserID       int                    `json:"user_id,omitempty"`
	Locality     string                 `json:"locality,omitempty"`
	Number       string                 `json:"number,omitempty"`
	Complement   string                 `json:"complement,omitempty"`
	Neighborhood string                 `json:"neighborhood,omitempty"`
	City         string                 `json:"city,omitempty"`
	State        address.BrazilianState `json:"state,omitempty"`
	CEP          string                 `json:"cep,omitempty"`
	Country      string                 `json:"country,omitempty"`
	CreatedAt    time.Time              `json:"created_at,omitempty"`
	UpdatedAt    time.Time              `json:"updated_at,omitempty"`
	DeletedAt    time.Time              `json:"deleted_at,omitempty"`
}

type addressInputDTO struct {
	ID           int                    `json:"id"`
	UserID       int                    `json:"user_id" binding:"required"`
	Locality     string                 `json:"locality" binding:"required"`
	Number       string                 `json:"number" binding:"required"`
	Complement   string                 `json:"complement,omitempty" binding:"required"`
	Neighborhood string                 `json:"neighborhood" binding:"required"`
	City         string                 `json:"city" binding:"required"`
	State        address.BrazilianState `json:"state" binding:"required"`
	CEP          string                 `json:"cep" binding:"required"`
	Country      string                 `json:"country" binding:"required"`
}

func mapAddressToOutputDTO(address address.Address) *addressOutputDTO {
	return &addressOutputDTO{
		ID:           address.ID,
		UserID:       address.UserID,
		Locality:     address.Locality,
		Number:       address.Number,
		Complement:   address.Complement,
		Neighborhood: address.Neighborhood,
		City:         address.City,
		State:        address.State,
		CEP:          address.CEP,
		Country:      address.Country,
		CreatedAt:    address.CreatedAt,
		UpdatedAt:    address.UpdatedAt,
		DeletedAt:    address.DeletedAt,
	}
}

func mapInputDTOToAddress(input addressInputDTO) *address.Address {
	return &address.Address{
		ID:           input.ID,
		UserID:       input.UserID,
		Locality:     input.Locality,
		Number:       input.Number,
		Complement:   input.Complement,
		Neighborhood: input.Neighborhood,
		City:         input.City,
		State:        input.State,
		CEP:          input.CEP,
		Country:      input.Country,
	}
}
