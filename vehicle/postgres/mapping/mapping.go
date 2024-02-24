package mapping

import (
	"github.com/LucasMateus-eng/operations-service/vehicle"
	vehicle_dto "github.com/LucasMateus-eng/operations-service/vehicle/postgres/dto"
)

func MapVehicleToDTO(vehicle *vehicle.Vehicle) *vehicle_dto.VehicleDTO {
	return &vehicle_dto.VehicleDTO{
		ID:                  vehicle.ID,
		Brand:               vehicle.Attributes.Brand,
		Model:               vehicle.Attributes.Model,
		YearOfManufacture:   vehicle.Attributes.YearOfManufacture,
		Plate:               vehicle.LegalInformation.Plate,
		Renavam:             vehicle.LegalInformation.Renavam,
		LicensingExpiryDate: vehicle.LegalInformation.Licensing.ExpiryDate,
		LicensingStatus:     vehicle.LegalInformation.Licensing.Status.String(),
		CreatedAt:           vehicle.CreatedAt,
		UpdatedAt:           vehicle.UpdatedAt,
		DeletedAt:           vehicle.DeletedAt,
	}
}

func MapDTOToVehicle(vehicleDTO *vehicle_dto.VehicleDTO) (*vehicle.Vehicle, error) {
	licensingStatus, err := vehicle.GetLicensingStatus(vehicleDTO.LicensingStatus)
	if err != nil {
		return nil, err
	}

	return &vehicle.Vehicle{
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
				Status:     *licensingStatus,
			},
		},
		CreatedAt: vehicleDTO.CreatedAt,
		UpdatedAt: vehicleDTO.UpdatedAt,
		DeletedAt: vehicleDTO.DeletedAt,
	}, nil
}
