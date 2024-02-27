package drivervehicle

import (
	"context"

	"github.com/LucasMateus-eng/operations-service/driver"
	"github.com/LucasMateus-eng/operations-service/internal/logging"
	"github.com/LucasMateus-eng/operations-service/vehicle"
)

type Service struct {
	repo   Repository
	logger *logging.Logging
}

func NewService(r Repository, l *logging.Logging) *Service {
	return &Service{
		repo:   r,
		logger: l,
	}
}

func (s *Service) GetByID(ctx context.Context, driverID, vehicleID int) (*DriverVehicle, error) {
	driverVehicle, err := s.repo.GetByID(ctx, driverID, vehicleID)
	if err != nil {
		return nil, err
	}

	return driverVehicle, nil
}

func (s *Service) GetDriverListByVehicleID(ctx context.Context, specification *DriverVehicleSpecification) (*[]driver.Driver, error) {
	drivers, err := s.repo.GetDriverListByVehicleID(ctx, specification)
	if err != nil {
		return nil, err
	}

	return drivers, nil
}

func (s *Service) GetVehicleListByDriverID(ctx context.Context, specification *DriverVehicleSpecification) (*[]vehicle.Vehicle, error) {
	vehicles, err := s.repo.GetVehicleListByDriverID(ctx, specification)
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (s *Service) Create(ctx context.Context, dv *DriverVehicle) (*DriverVehicle, error) {
	driverVehicle, err := s.repo.Create(ctx, dv)
	if err != nil {
		return nil, err
	}

	return driverVehicle, nil
}

func (s *Service) Delete(ctx context.Context, driverID, vehicleID int) error {
	return s.repo.Delete(ctx, driverID, vehicleID)
}
