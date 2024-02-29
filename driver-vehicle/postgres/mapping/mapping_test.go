package mapping

import (
	"testing"
	"time"

	driver_vehicle "github.com/LucasMateus-eng/operations-service/driver-vehicle"
	driver_vehicle_dto "github.com/LucasMateus-eng/operations-service/driver-vehicle/postgres/dto"
	driver_dto "github.com/LucasMateus-eng/operations-service/driver/postgres/dto"
	vehicle_dto "github.com/LucasMateus-eng/operations-service/vehicle/postgres/dto"
	"github.com/go-playground/assert/v2"
)

var (
	mockedTime = time.Now()
)

func TestMapDriverVehicleToDTO(t *testing.T) {
	driverVehicle := &driver_vehicle.DriverVehicle{
		DriverID:  1,
		VehicleID: 1,
		CreatedAt: mockedTime,
		UpdatedAt: mockedTime,
		DeletedAt: mockedTime,
	}

	expectedDTO := &driver_vehicle_dto.DriverVehicleDTO{
		DriverID:  driverVehicle.DriverID,
		Driver:    driver_dto.DriverDTO{},
		VehicleID: driverVehicle.VehicleID,
		Vehicle:   vehicle_dto.VehicleDTO{},
		CreatedAt: driverVehicle.CreatedAt,
		UpdatedAt: driverVehicle.UpdatedAt,
		DeletedAt: driverVehicle.DeletedAt,
	}

	actualDTO := MapDriverVehicleToDTO(driverVehicle)
	assert.Equal(t, expectedDTO, actualDTO)
}

func TestMapDTOToDriverVehicle(t *testing.T) {
	driverVehicleDTO := &driver_vehicle_dto.DriverVehicleDTO{
		DriverID:  1,
		VehicleID: 1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}

	expectedDriverVehicle := &driver_vehicle.DriverVehicle{
		DriverID:  driverVehicleDTO.DriverID,
		VehicleID: driverVehicleDTO.VehicleID,
		CreatedAt: driverVehicleDTO.CreatedAt,
		UpdatedAt: driverVehicleDTO.UpdatedAt,
		DeletedAt: driverVehicleDTO.DeletedAt,
	}

	actualDriverVehicle := MapDTOToDriverVehicle(driverVehicleDTO)
	assert.Equal(t, expectedDriverVehicle, actualDriverVehicle)
}
