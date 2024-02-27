package gin

import (
	"time"

	drivervehicle "github.com/LucasMateus-eng/operations-service/driver-vehicle"
)

type driverVehicleOutputDTO struct {
	DriverID  int       `json:"driver_id"`
	VehicleID int       `json:"vehicle_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type driverVehicleInputDTO struct {
	DriverID  int `json:"driver_id" binding:"required"`
	VehicleID int `json:"vehicle_id" binding:"required"`
}

func mapDriverVehicleToOutputDTO(driverVehicle *drivervehicle.DriverVehicle) *driverVehicleOutputDTO {
	return &driverVehicleOutputDTO{
		DriverID:  driverVehicle.DriverID,
		VehicleID: driverVehicle.VehicleID,
		CreatedAt: driverVehicle.CreatedAt,
		UpdatedAt: driverVehicle.UpdatedAt,
		DeletedAt: driverVehicle.DeletedAt,
	}
}

func mapInputDTOToDriverVehicle(input *driverVehicleInputDTO) *drivervehicle.DriverVehicle {
	return &drivervehicle.DriverVehicle{
		DriverID:  input.DriverID,
		VehicleID: input.VehicleID,
	}
}
