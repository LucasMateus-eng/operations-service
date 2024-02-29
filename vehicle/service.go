package vehicle

import (
	"context"

	"github.com/LucasMateus-eng/operations-service/internal/logging"
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

func (s *Service) GetByID(ctx context.Context, id int64) (*Vehicle, error) {
	s.logger.Debug("[VEHICLE] GetByID - DEBUG: ", map[string]any{
		"vehicleID": id,
	})
	vehicle, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("[VEHICLE] GetByID - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return vehicle, nil
}

func (s *Service) GetByPlate(ctx context.Context, plate string) (*Vehicle, error) {
	s.logger.Debug("[VEHICLE] GetByPlate - DEBUG: ", map[string]any{
		"vehiclePlate": plate,
	})
	vehicle, err := s.repo.GetByPlate(ctx, plate)
	if err != nil {
		s.logger.Error("[VEHICLE] GetByPlate - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return vehicle, nil
}

func (s *Service) GetByRenavam(ctx context.Context, renavam string) (*Vehicle, error) {
	s.logger.Debug("[VEHICLE] GetByRenavam - DEBUG: ", map[string]any{
		"vehicleRenavam": renavam,
	})
	vehicle, err := s.repo.GetByRenavam(ctx, renavam)
	if err != nil {
		s.logger.Error("[VEHICLE] GetByRenavam - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return vehicle, nil
}

func (s *Service) List(ctx context.Context, specification *VehicleSpectification) (*[]Vehicle, error) {
	s.logger.Debug("[VEHICLE] List - DEBUG: ", map[string]any{
		"specification": specification,
	})
	vehicles, err := s.repo.List(ctx, specification)
	if err != nil {
		s.logger.Error("[VEHICLE] List - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return vehicles, nil
}

func (s *Service) Create(ctx context.Context, v *Vehicle) (int64, error) {
	s.logger.Debug("[VEHICLE] Create - DEBUG: ", map[string]any{
		"vehicle": v,
	})
	vehicleID, err := s.repo.Create(ctx, v)
	if err != nil {
		s.logger.Error("[VEHICLE] Create - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return 0, err
	}

	return vehicleID, nil
}

func (s *Service) Update(ctx context.Context, v *Vehicle) error {
	s.logger.Debug("[VEHICLE] Update - DEBUG: ", map[string]any{
		"vehicle": v,
	})
	err := s.repo.Update(ctx, v)
	if err != nil {
		s.logger.Error("[VEHICLE] Update - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return err
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	s.logger.Debug("[VEHICLE] Delete - DEBUG: ", map[string]any{
		"vehicleID": id,
	})
	err := s.repo.Delete(ctx, id)
	if err != nil {
		s.logger.Error("[VEHICLE] Delete - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return err
	}

	return nil
}
