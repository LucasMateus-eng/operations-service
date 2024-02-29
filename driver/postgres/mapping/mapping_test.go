package mapping

import (
	"testing"
	"time"

	"github.com/LucasMateus-eng/operations-service/address"
	address_dto "github.com/LucasMateus-eng/operations-service/address/postgres/dto"
	"github.com/LucasMateus-eng/operations-service/driver"
	driver_dto "github.com/LucasMateus-eng/operations-service/driver/postgres/dto"
	user_dto "github.com/LucasMateus-eng/operations-service/user/postgres/dto"
	"github.com/LucasMateus-eng/operations-service/vehicle"
	vehicle_dto "github.com/LucasMateus-eng/operations-service/vehicle/postgres/dto"
	"github.com/go-playground/assert/v2"
)

var (
	mockedTime = time.Now()
)

func TestMapDriverToDTO(t *testing.T) {
	driver := &driver.Driver{
		ID:     1,
		UserID: 1,
		Attributes: driver.DriverAttributes{
			Name:        "John Doe",
			DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		LegalInformation: driver.DriverLegalInformation{
			RG:            "123456",
			CPF:           "7891011",
			DriverLicense: "DL123",
		},
		Contact: driver.Contact{
			CellPhone: "123456789",
			Email:     "john@example.com",
		},
		CreatedAt: mockedTime,
		UpdatedAt: mockedTime,
		DeletedAt: mockedTime,
	}

	expectedDTO := &driver_dto.DriverDTO{
		ID:            driver.ID,
		Name:          driver.Attributes.Name,
		RG:            driver.LegalInformation.RG,
		CPF:           driver.LegalInformation.CPF,
		DriverLicense: driver.LegalInformation.DriverLicense,
		DateOfBirth:   driver.Attributes.DateOfBirth,
		CellPhone:     driver.Contact.CellPhone,
		Email:         driver.Contact.Email,
		UserID:        driver.UserID,
		Vehicles:      []vehicle_dto.VehicleDTO{},
		CreatedAt:     driver.CreatedAt,
		UpdatedAt:     driver.UpdatedAt,
		DeletedAt:     driver.DeletedAt,
	}

	actualDTO := MapDriverToDTO(driver)
	assert.Equal(t, expectedDTO, actualDTO)
}

func TestMapDTOToDriver(t *testing.T) {
	userAddressDTO := address_dto.AddressDTO{
		ID:           1,
		UserID:       1,
		Locality:     "Main Street",
		Number:       "123",
		Complement:   "Apt 2",
		Neighborhood: "Downtown",
		City:         "City",
		State:        "ACRE",
		CEP:          "12345678",
		Country:      "Country",
		CreatedAt:    mockedTime,
		UpdatedAt:    mockedTime,
		DeletedAt:    mockedTime,
	}

	userDTO := user_dto.UserDTO{
		ID:         1,
		Username:   "john_doe",
		Role:       "driver",
		AddressDTO: &userAddressDTO,
	}

	vehicleDTO := vehicle_dto.VehicleDTO{
		ID:                  1,
		Brand:               "Toyota",
		Model:               "Corolla",
		YearOfManufacture:   mockedTime,
		Plate:               "ABC1234",
		Renavam:             "123456789",
		LicensingExpiryDate: mockedTime,
		LicensingStatus:     "REGULAR",
		CreatedAt:           mockedTime,
		UpdatedAt:           mockedTime,
		DeletedAt:           mockedTime,
	}

	vehicleDTOs := []vehicle_dto.VehicleDTO{vehicleDTO}

	driverDTO := &driver_dto.DriverDTO{
		ID:            1,
		Name:          "John Doe",
		RG:            "123456",
		CPF:           "7891011",
		DriverLicense: "DL123",
		DateOfBirth:   time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		CellPhone:     "123456789",
		Email:         "john@example.com",
		UserID:        1,
		User:          userDTO,
		Vehicles:      vehicleDTOs,
		CreatedAt:     mockedTime,
		UpdatedAt:     mockedTime,
		DeletedAt:     mockedTime,
	}

	userAddressDTOWithInvalidState := address_dto.AddressDTO{
		ID:           1,
		UserID:       1,
		Locality:     "Main Street",
		Number:       "123",
		Complement:   "Apt 2",
		Neighborhood: "Downtown",
		City:         "City",
		State:        "MISSOURI",
		CEP:          "12345678",
		Country:      "Country",
		CreatedAt:    mockedTime,
		UpdatedAt:    mockedTime,
		DeletedAt:    mockedTime,
	}

	userDTOWithInvalidAddress := user_dto.UserDTO{
		ID:         1,
		Username:   "john_doe",
		Role:       "driver",
		AddressDTO: &userAddressDTOWithInvalidState,
	}

	driverDTOWithInvalidUser := &driver_dto.DriverDTO{
		ID:            1,
		Name:          "John Doe",
		RG:            "123456",
		CPF:           "7891011",
		DriverLicense: "DL123",
		DateOfBirth:   time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		CellPhone:     "123456789",
		Email:         "john@example.com",
		UserID:        1,
		User:          userDTOWithInvalidAddress,
		Vehicles:      vehicleDTOs,
		CreatedAt:     mockedTime,
		UpdatedAt:     mockedTime,
		DeletedAt:     mockedTime,
	}

	expectedAddress := &address.Address{
		ID:           userAddressDTO.ID,
		UserID:       userAddressDTO.UserID,
		Locality:     userAddressDTO.Locality,
		Number:       userAddressDTO.Number,
		Complement:   userAddressDTO.Complement,
		Neighborhood: userAddressDTO.Neighborhood,
		City:         userAddressDTO.City,
		State:        address.AC,
		CEP:          userAddressDTO.CEP,
		Country:      userAddressDTO.Country,
		CreatedAt:    userAddressDTO.CreatedAt,
		UpdatedAt:    userAddressDTO.UpdatedAt,
		DeletedAt:    userAddressDTO.DeletedAt,
	}

	expectedVehicle := vehicle.Vehicle{
		ID: vehicleDTO.ID,
		Attributes: vehicle.VehicleAttributes{
			Brand:             vehicleDTO.Brand,
			Model:             vehicleDTO.Model,
			YearOfManufacture: vehicleDTO.YearOfManufacture,
		},
		LegalInformation: vehicle.VehicleLegalInformation{
			Plate:   vehicleDTO.Plate,
			Renavam: vehicleDTO.Renavam,
			Licensing: vehicle.Licensing{
				ExpiryDate: vehicleDTO.LicensingExpiryDate,
				Status:     vehicle.REGULAR,
			},
		},
		CreatedAt: vehicleDTO.CreatedAt,
		UpdatedAt: vehicleDTO.UpdatedAt,
		DeletedAt: vehicleDTO.DeletedAt,
	}

	expectedDriver := &driver.Driver{
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
		Address: expectedAddress,
		Contact: driver.Contact{
			CellPhone: driverDTO.CellPhone,
			Email:     driverDTO.Email,
		},
		Vehicles:  []vehicle.Vehicle{expectedVehicle},
		CreatedAt: driverDTO.CreatedAt,
		UpdatedAt: driverDTO.UpdatedAt,
		DeletedAt: driverDTO.DeletedAt,
	}

	tests := []struct {
		name    string
		arg     *driver_dto.DriverDTO
		want    *driver.Driver
		wantErr bool
	}{
		{
			name:    "Dado um DTO de Driver quando a função de mapeamento é chamada então a conversão é um sucesso",
			arg:     driverDTO,
			want:    expectedDriver,
			wantErr: false,
		},
		{
			name:    "Dado um DTO de Driver com um DTO de Address com um estado inválido quando a função de mapeamento é chamada então a conversão falha",
			arg:     driverDTOWithInvalidUser,
			want:    nil,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			actualDriver, err := MapDTOToDriver(test.arg)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualDriver)
		})
	}
}
