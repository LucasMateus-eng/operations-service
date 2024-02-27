package mapping

import (
	"github.com/LucasMateus-eng/operations-service/address"
	"github.com/LucasMateus-eng/operations-service/driver"
	driver_dto "github.com/LucasMateus-eng/operations-service/driver/postgres/dto"
	"github.com/LucasMateus-eng/operations-service/vehicle"
	vehicle_dto "github.com/LucasMateus-eng/operations-service/vehicle/postgres/dto"
	vehicle_mapping "github.com/LucasMateus-eng/operations-service/vehicle/postgres/mapping"
)

func MapDriverToDTO(driver *driver.Driver) *driver_dto.DriverDTO {
	vehicleDTOs := make([]vehicle_dto.VehicleDTO, len(driver.Vehicles))
	for i, vehicle := range driver.Vehicles {
		vehicleDTOs[i] = *vehicle_mapping.MapVehicleToDTO(&vehicle)
	}

	return &driver_dto.DriverDTO{
		ID:            driver.ID,
		Name:          driver.Attributes.Name,
		RG:            driver.LegalInformation.RG,
		CPF:           driver.LegalInformation.CPF,
		DriverLicense: driver.LegalInformation.DriverLicense,
		DateOfBirth:   driver.Attributes.DateOfBirth,
		CellPhone:     driver.Contact.CellPhone,
		Email:         driver.Contact.Email,
		UserID:        driver.UserID,
		Vehicles:      vehicleDTOs,
		CreatedAt:     driver.CreatedAt,
		UpdatedAt:     driver.UpdatedAt,
		DeletedAt:     driver.DeletedAt,
	}
}

func MapDTOToDriver(driverDTO *driver_dto.DriverDTO) (*driver.Driver, error) {
	vehicles, err := mapVehicles(driverDTO.Vehicles)
	if err != nil {
		return nil, err
	}

	state, err := address.GetBrazilianState(driverDTO.User.AddressDTO.State)
	if err != nil {
		return nil, err
	}

	return &driver.Driver{
		ID:     driverDTO.ID,
		UserID: driverDTO.UserID,
		Attributes: driver.DriverAttributes{
			Name:        driverDTO.Name,
			DateOfBirth: driverDTO.DateOfBirth,
		},
		LegalInformation: driver.DriverLegalInformation{
			RG:            driverDTO.RG,
			CPF:           driverDTO.CPF,
			DriverLicense: driverDTO.DriverLicense,
		},
		Address: &address.Address{
			ID:           driverDTO.User.AddressDTO.ID,
			UserID:       driverDTO.User.AddressDTO.UserID,
			Locality:     driverDTO.User.AddressDTO.Locality,
			Number:       driverDTO.User.AddressDTO.Number,
			Complement:   driverDTO.User.AddressDTO.Complement,
			Neighborhood: driverDTO.User.AddressDTO.Neighborhood,
			City:         driverDTO.User.AddressDTO.City,
			State:        state,
			CEP:          driverDTO.User.AddressDTO.CEP,
			Country:      driverDTO.User.AddressDTO.Country,
			CreatedAt:    driverDTO.User.AddressDTO.CreatedAt,
			UpdatedAt:    driverDTO.User.AddressDTO.UpdatedAt,
			DeletedAt:    driverDTO.User.AddressDTO.DeletedAt,
		},
		Contact: driver.Contact{
			CellPhone: driverDTO.CellPhone,
			Email:     driverDTO.Email,
		},
		Vehicles:  vehicles,
		CreatedAt: driverDTO.CreatedAt,
		UpdatedAt: driverDTO.UpdatedAt,
		DeletedAt: driverDTO.DeletedAt,
	}, nil
}

func mapVehicles(vehicleDTOs []vehicle_dto.VehicleDTO) ([]vehicle.Vehicle, error) {
	vehicles := make([]vehicle.Vehicle, len(vehicleDTOs))
	for i, vehicleDTO := range vehicleDTOs {
		v, err := vehicle_mapping.MapDTOToVehicle(&vehicleDTO)
		if err != nil {
			return nil, err
		}
		vehicles[i] = *v
	}
	return vehicles, nil
}
