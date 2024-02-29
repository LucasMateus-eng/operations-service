package mapping

import (
	"testing"
	"time"

	"github.com/LucasMateus-eng/operations-service/vehicle"
	vehicle_dto "github.com/LucasMateus-eng/operations-service/vehicle/postgres/dto"
	"github.com/go-playground/assert/v2"
)

var (
	mockedTime = time.Now()
)

func TestMapVehicleToDTO(t *testing.T) {
	vehicle := &vehicle.Vehicle{
		ID: 1,
		Attributes: vehicle.VehicleAttributes{
			Brand:             "Toyota",
			Model:             "Corolla",
			YearOfManufacture: mockedTime,
		},
		LegalInformation: vehicle.VehicleLegalInformation{
			Plate:   "ABC1234",
			Renavam: "123456789",
			Licensing: vehicle.Licensing{
				ExpiryDate: mockedTime,
				Status:     vehicle.REGULAR,
			},
		},
		CreatedAt: mockedTime,
		UpdatedAt: mockedTime,
		DeletedAt: mockedTime,
	}

	expectedDTO := &vehicle_dto.VehicleDTO{
		ID:                  1,
		Brand:               "Toyota",
		Model:               "Corolla",
		YearOfManufacture:   vehicle.Attributes.YearOfManufacture,
		Plate:               "ABC1234",
		Renavam:             "123456789",
		LicensingExpiryDate: vehicle.LegalInformation.Licensing.ExpiryDate,
		LicensingStatus:     "REGULAR",
		CreatedAt:           vehicle.CreatedAt,
		UpdatedAt:           vehicle.UpdatedAt,
		DeletedAt:           vehicle.DeletedAt,
	}

	actualDTO := MapVehicleToDTO(vehicle)
	assert.Equal(t, expectedDTO, actualDTO)
}

func TestMapDTOToVehicle(t *testing.T) {
	vehicleDTO := &vehicle_dto.VehicleDTO{
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

	vehicleDTOWithInvalidLicensingStatus := &vehicle_dto.VehicleDTO{
		ID:                  1,
		Brand:               "Toyota",
		Model:               "Corolla",
		YearOfManufacture:   mockedTime,
		Plate:               "ABC1234",
		Renavam:             "123456789",
		LicensingExpiryDate: mockedTime,
		LicensingStatus:     "OPTIMIZED",
		CreatedAt:           mockedTime,
		UpdatedAt:           mockedTime,
		DeletedAt:           mockedTime,
	}

	expectedVehicle := &vehicle.Vehicle{
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

	tests := []struct {
		name    string
		arg     *vehicle_dto.VehicleDTO
		want    *vehicle.Vehicle
		wantErr bool
	}{
		{
			name:    "Dado um DTO de Vehicle quando a função de mapeamento é chamada então a conversão é um sucesso",
			arg:     vehicleDTO,
			want:    expectedVehicle,
			wantErr: false,
		},
		{
			name:    "Dado um DTO de Vehicle com um estado de licenciamento inválido quando a função de mapeamento é chamada então a conversão falha",
			arg:     vehicleDTOWithInvalidLicensingStatus,
			want:    nil,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			actualVehicle, err := MapDTOToVehicle(test.arg)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualVehicle)
		})
	}
}
