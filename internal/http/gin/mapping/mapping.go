package mapping

import (
	"github.com/LucasMateus-eng/operations-service/address"
	"github.com/LucasMateus-eng/operations-service/driver"
	drivervehicle "github.com/LucasMateus-eng/operations-service/driver-vehicle"
	gin_dto "github.com/LucasMateus-eng/operations-service/internal/http/gin/dto"
	"github.com/LucasMateus-eng/operations-service/user"
	"github.com/LucasMateus-eng/operations-service/vehicle"
)

func MapAddressToOutputDTO(address address.Address) *gin_dto.AddressOutputDTO {
	return &gin_dto.AddressOutputDTO{
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

func MapDriverVehicleToOutputDTO(driverVehicle drivervehicle.DriverVehicle) *gin_dto.DriverVehicleOutputDTO {
	return &gin_dto.DriverVehicleOutputDTO{
		DriverID:  driverVehicle.DriverID,
		VehicleID: driverVehicle.VehicleID,
		CreatedAt: driverVehicle.CreatedAt,
		UpdatedAt: driverVehicle.UpdatedAt,
		DeletedAt: driverVehicle.DeletedAt,
	}
}

func MapInputDTOToDriverVehicle(input gin_dto.DriverVehicleInputDTO) *drivervehicle.DriverVehicle {
	return &drivervehicle.DriverVehicle{
		DriverID:  input.DriverID,
		VehicleID: input.VehicleID,
	}
}

func MapDriverToOutputDTO(driver driver.Driver) *gin_dto.DriverOutputDTO {
	return &gin_dto.DriverOutputDTO{
		ID:            driver.ID,
		UserID:        driver.UserID,
		Name:          driver.Attributes.Name,
		DateOfBirth:   driver.Attributes.DateOfBirth,
		RG:            driver.LegalInformation.RG,
		CPF:           driver.LegalInformation.CPF,
		DriverLicense: driver.LegalInformation.DriverLicense,
		CellPhone:     driver.Contact.CellPhone,
		Email:         driver.Contact.Email,
		Address:       MapAddressToOutputDTO(*driver.Address),
		Vehicles:      MapVehicleListToOutputDTO(driver.Vehicles),
		CreatedAt:     driver.CreatedAt,
		UpdatedAt:     driver.UpdatedAt,
		DeletedAt:     driver.DeletedAt,
	}
}

func MapInputDTOToDriver(input gin_dto.DriverInputDTO) *driver.Driver {
	return &driver.Driver{
		ID: input.ID,
		Attributes: driver.DriverAttributes{
			Name:        input.Name,
			DateOfBirth: input.DateOfBirth,
		},
		LegalInformation: driver.DriverLegalInformation{
			RG:            input.RG,
			CPF:           input.CPF,
			DriverLicense: input.DriverLicense,
		},
		Contact: driver.Contact{
			CellPhone: input.CellPhone,
			Email:     input.Email,
		},
	}
}

func MapInputDTOToDriverSpecification(input gin_dto.DriverSpecificationInputDTO) *driver.DriverSpecification {
	return &driver.DriverSpecification{
		Page:     input.Page,
		PageSize: input.PageSize,
	}
}

func MapUserToOutputDTO(user user.User) *gin_dto.UserOutputDTO {
	return &gin_dto.UserOutputDTO{
		ID:             user.ID,
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		Role:           user.Role,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		DeletedAt:      user.DeletedAt,
	}
}

func MapInputDTOToUser(input gin_dto.UserInputDTO) *user.User {
	return &user.User{
		ID:             input.ID,
		Username:       input.Username,
		HashedPassword: input.HashedPassword,
		Role:           input.Role,
	}
}

func MapInputDTOToVehicleSpecification(input gin_dto.VehicleSpecificationInputDTO) *vehicle.VehicleSpectification {
	return &vehicle.VehicleSpectification{
		Attributes: vehicle.VehicleAttributes{
			Brand:             input.Brand,
			Model:             input.Model,
			YearOfManufacture: input.YearOfManufacture,
		},
		Licensing: vehicle.Licensing{
			ExpiryDate: input.LicensingExpiryDate,
			Status:     input.LicensingStatus,
		},
		Page:     input.Page,
		PageSize: input.PageSize,
	}
}

func MapVehicleToOutputDTO(vehicle vehicle.Vehicle) *gin_dto.VehicleOutputDTO {
	return &gin_dto.VehicleOutputDTO{
		ID:                  vehicle.ID,
		Brand:               vehicle.Attributes.Brand,
		Model:               vehicle.Attributes.Model,
		YearOfManufacture:   vehicle.Attributes.YearOfManufacture,
		Plate:               vehicle.LegalInformation.Plate,
		Renavam:             vehicle.LegalInformation.Renavam,
		LicensingExpiryDate: vehicle.LegalInformation.Licensing.ExpiryDate,
		LicensingStatus:     vehicle.LegalInformation.Licensing.Status,
		CreatedAt:           vehicle.CreatedAt,
		UpdatedAt:           vehicle.UpdatedAt,
		DeletedAt:           vehicle.DeletedAt,
	}
}

func MapInputDTOToVehicle(input gin_dto.VehicleInputDTO) *vehicle.Vehicle {
	return &vehicle.Vehicle{
		ID: input.ID,
		Attributes: vehicle.VehicleAttributes{
			Brand:             input.Brand,
			Model:             input.Model,
			YearOfManufacture: input.YearOfManufacture,
		},
		LegalInformation: vehicle.VehicleLegalInformation{
			Plate:   input.Plate,
			Renavam: input.Renavam,
			Licensing: vehicle.Licensing{
				ExpiryDate: input.LicensingExpiryDate,
				Status:     input.LicensingStatus,
			},
		},
	}
}

func MapVehicleListToOutputDTO(vehicles []vehicle.Vehicle) []gin_dto.VehicleOutputDTO {
	vehicleDTOs := make([]gin_dto.VehicleOutputDTO, 0, len(vehicles))
	for i, v := range vehicles {
		vehicleDTOs[i] = *MapVehicleToOutputDTO(v)
	}
	return vehicleDTOs
}
