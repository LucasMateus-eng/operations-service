package gin

import (
	"time"

	"github.com/LucasMateus-eng/operations-service/vehicle"
)

type vehicleOutputDTO struct {
	ID                int                     `json:"id"`
	Brand             string                  `json:"brand,omitempty"`
	Model             string                  `json:"model,omitempty"`
	YearOfManufacture time.Time               `json:"year_of_manufacture,omitempty"`
	Plate             string                  `json:"plate,omitempty"`
	Renavam           string                  `json:"renavam,omitempty"`
	LicensingExpiry   time.Time               `json:"licensing_expiry_date,omitempty"`
	LicensingStatus   vehicle.LicensingStatus `json:"licensing_status,omitempty"`
	CreatedAt         time.Time               `json:"created_at,omitempty"`
	UpdatedAt         time.Time               `json:"updated_at,omitempty"`
	DeletedAt         time.Time               `json:"deleted_at,omitempty"`
}

type vehicleInputDTO struct {
	ID                int                     `json:"id" binding:"required"`
	Brand             string                  `json:"brand" binding:"required"`
	Model             string                  `json:"model" binding:"required"`
	YearOfManufacture time.Time               `json:"year_of_manufacture" binding:"required"`
	Plate             string                  `json:"plate" binding:"required"`
	Renavam           string                  `json:"renavam" binding:"required"`
	LicensingExpiry   time.Time               `json:"licensing_expiry_date" binding:"required"`
	LicensingStatus   vehicle.LicensingStatus `json:"licensing_status" binding:"required"`
}

func mapVehicleToOutputDTO(vehicle *vehicle.Vehicle) *vehicleOutputDTO {
	return &vehicleOutputDTO{
		ID:                vehicle.ID,
		Brand:             vehicle.Attributes.Brand,
		Model:             vehicle.Attributes.Model,
		YearOfManufacture: vehicle.Attributes.YearOfManufacture,
		Plate:             vehicle.LegalInformation.Plate,
		Renavam:           vehicle.LegalInformation.Renavam,
		LicensingExpiry:   vehicle.LegalInformation.Licensing.ExpiryDate,
		LicensingStatus:   vehicle.LegalInformation.Licensing.Status,
		CreatedAt:         vehicle.CreatedAt,
		UpdatedAt:         vehicle.UpdatedAt,
		DeletedAt:         vehicle.DeletedAt,
	}
}

func mapInputDTOToVehicle(input *vehicleOutputDTO) *vehicle.Vehicle {
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
				ExpiryDate: input.LicensingExpiry,
				Status:     input.LicensingStatus,
			},
		},
	}
}

func mapVehicleListToOutputDTO(vehicles []vehicle.Vehicle) []vehicleOutputDTO {
	vehicleDTOs := make([]vehicleOutputDTO, 0, len(vehicles))
	for i, v := range vehicles {
		vehicleDTOs[i] = *mapVehicleToOutputDTO(&v)
	}
	return vehicleDTOs
}
