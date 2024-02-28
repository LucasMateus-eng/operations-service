package drivervehicle

import (
	"context"
	"time"

	"github.com/LucasMateus-eng/operations-service/driver"
	"github.com/LucasMateus-eng/operations-service/vehicle"
)

type DriverVehicle struct {
	DriverID  int64
	VehicleID int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type DriverVehicleSpecification struct {
	VehicleID, DriverID int64
	Page, PageSize      int
}

type Reading interface {
	GetByID(ctx context.Context, driverID, vehicleID int64) (*DriverVehicle, error)
	GetDriverListByVehicleID(ctx context.Context, specification *DriverVehicleSpecification) (*[]driver.Driver, error)
	GetVehicleListByDriverID(ctx context.Context, specification *DriverVehicleSpecification) (*[]vehicle.Vehicle, error)
}

type Writing interface {
	Create(ctx context.Context, dv *DriverVehicle) (*DriverVehicle, error)
	Delete(ctx context.Context, driverID, vehicleID int64) error
}

type Repository interface {
	Reading
	Writing
}

type UseCase interface {
	GetByID(ctx context.Context, driverID, vehicleID int64) (*DriverVehicle, error)
	GetDriverListByVehicleID(ctx context.Context, specification *DriverVehicleSpecification) (*[]driver.DriverList, error)
	GetVehicleListByDriverID(ctx context.Context, specification *DriverVehicleSpecification) (*[]vehicle.VehicleList, error)
	Create(ctx context.Context, dv *DriverVehicle) (int64, error)
	Delete(ctx context.Context, driverID, vehicleID int64) error
}
