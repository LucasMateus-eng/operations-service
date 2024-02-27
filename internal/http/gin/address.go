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
