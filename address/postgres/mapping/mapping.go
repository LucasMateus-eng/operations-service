package mapping

import (
	"github.com/LucasMateus-eng/operations-service/address"
	address_dto "github.com/LucasMateus-eng/operations-service/address/postgres/dto"
)

func MapAddressToDTO(address *address.Address) *address_dto.AddressDTO {
	return &address_dto.AddressDTO{
		ID:           address.ID,
		Locality:     address.Locality,
		Number:       address.Number,
		Complement:   address.Complement,
		Neighborhood: address.Neighborhood,
		City:         address.City,
		State:        address.State.String(),
		CEP:          address.CEP,
		Country:      address.Country,
		UserID:       address.UserID,
		CreatedAt:    address.CreatedAt,
		UpdatedAt:    address.UpdatedAt,
		DeletedAt:    address.DeletedAt,
	}
}

func MapDTOToAddress(addressDTO *address_dto.AddressDTO) (*address.Address, error) {
	state, err := address.GetBrazilianState(addressDTO.State)
	if err != nil {
		return nil, err
	}

	return &address.Address{
		ID:           addressDTO.ID,
		Locality:     addressDTO.Locality,
		Number:       addressDTO.Number,
		Complement:   addressDTO.Complement,
		Neighborhood: addressDTO.Neighborhood,
		City:         addressDTO.City,
		State:        state,
		CEP:          addressDTO.CEP,
		Country:      addressDTO.Country,
		UserID:       addressDTO.UserID,
		CreatedAt:    addressDTO.CreatedAt,
		UpdatedAt:    addressDTO.UpdatedAt,
		DeletedAt:    addressDTO.DeletedAt,
	}, nil
}
