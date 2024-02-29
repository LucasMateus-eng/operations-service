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

func (s *Service) GetByID(ctx context.Context, driverID, vehicleID int64) (*DriverVehicle, error) {
	s.logger.Debug("[DRIVER-VEHICLE] GetByID - DEBUG: ", map[string]any{
		"driverID":  driverID,
		"vehicleID": vehicleID,
	})
	driverVehicle, err := s.repo.GetByID(ctx, driverID, vehicleID)
	if err != nil {
		s.logger.Error("[DRIVER-VEHICLE] GetByID - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return driverVehicle, nil
}

func (s *Service) GetDriverListByVehicleID(ctx context.Context, specification *DriverVehicleSpecification) (*[]driver.Driver, error) {
	s.logger.Debug("[DRIVER-VEHICLE] GetDriverListByVehicleID - DEBUG: ", map[string]any{
		"specification": specification,
	})
	drivers, err := s.repo.GetDriverListByVehicleID(ctx, specification)
	if err != nil {
		s.logger.Error("[DRIVER-VEHICLE] GetDriverListByVehicleID - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return drivers, nil
}

func (s *Service) GetVehicleListByDriverID(ctx context.Context, specification *DriverVehicleSpecification) (*[]vehicle.Vehicle, error) {
	s.logger.Debug("[DRIVER-VEHICLE] GetVehicleListByDriverID - DEBUG: ", map[string]any{
		"specification": specification,
	})
	vehicles, err := s.repo.GetVehicleListByDriverID(ctx, specification)
	if err != nil {
		s.logger.Error("[DRIVER-VEHICLE] GetVehicleListByDriverID - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return vehicles, nil
}

func (s *Service) Create(ctx context.Context, dv *DriverVehicle) (*DriverVehicle, error) {
	s.logger.Debug("[DRIVER-VEHICLE] Create - DEBUG: ", map[string]any{
		"driverVehicle": dv,
	})
	driverVehicle, err := s.repo.Create(ctx, dv)
	if err != nil {
		s.logger.Error("[DRIVER-VEHICLE] Create - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return driverVehicle, nil
}

func (s *Service) Delete(ctx context.Context, driverID, vehicleID int64) error {
	s.logger.Debug("[DRIVER-VEHICLE] Delete - DEBUG: ", map[string]any{
		"driverID":  driverID,
		"vehicleID": vehicleID,
	})
	err := s.repo.Delete(ctx, driverID, vehicleID)
	if err != nil {
		s.logger.Error("[DRIVER-VEHICLE] Delete - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return err
	}

	return nil
}
