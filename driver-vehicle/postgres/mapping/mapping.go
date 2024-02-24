package mapping

import (
	driver_vehicle "github.com/LucasMateus-eng/operations-service/driver-vehicle"
	driver_vehicle_dto "github.com/LucasMateus-eng/operations-service/driver-vehicle/postgres/dto"
	driver_dto "github.com/LucasMateus-eng/operations-service/driver/postgres/dto"
	vehicle_dto "github.com/LucasMateus-eng/operations-service/vehicle/postgres/dto"
)

func MapDriverVehicleToDTO(driverVehicle *driver_vehicle.DriverVehicle) *driver_vehicle_dto.DriverVehicleDTO {
	return &driver_vehicle_dto.DriverVehicleDTO{
		DriverID:  driverVehicle.DriverID,
		Driver:    driver_dto.DriverDTO{},
		VehicleID: driverVehicle.VehicleID,
		Vehicle:   vehicle_dto.VehicleDTO{},
		CreatedAt: driverVehicle.CreatedAt,
		UpdatedAt: driverVehicle.UpdatedAt,
		DeletedAt: driverVehicle.DeletedAt,
	}
}

func MapDTOToDriverVehicle(driverVehicleDTO *driver_vehicle_dto.DriverVehicleDTO) *driver_vehicle.DriverVehicle {
	return &driver_vehicle.DriverVehicle{
		DriverID:  driverVehicleDTO.DriverID,
		VehicleID: driverVehicleDTO.VehicleID,
		CreatedAt: driverVehicleDTO.CreatedAt,
		UpdatedAt: driverVehicleDTO.UpdatedAt,
		DeletedAt: driverVehicleDTO.DeletedAt,
	}
}
