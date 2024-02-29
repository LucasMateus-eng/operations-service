package driver

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

func (s *Service) GetByDriverId(ctx context.Context, id int64) (*Driver, error) {
	s.logger.Debug("[DRIVER] GetByDriverId - DEBUG: ", map[string]any{
		"driverID": id,
	})
	driver, err := s.repo.GetById(ctx, id)
	if err != nil {
		s.logger.Error("[DRIVER] GetByDriverId - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return driver, nil
}

func (s *Service) GetByUserId(ctx context.Context, userId int64) (*Driver, error) {
	s.logger.Debug("[DRIVER] GetByUserId - DEBUG: ", map[string]any{
		"userID": userId,
	})
	driver, err := s.repo.GetByUserId(ctx, userId)
	if err != nil {
		s.logger.Error("[DRIVER] GetByUserId - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return driver, nil
}

func (s *Service) GetByIdWithEagerLoading(ctx context.Context, id int64) (*Driver, error) {
	s.logger.Debug("[DRIVER] GetByIdWithEagerLoading - DEBUG: ", map[string]any{
		"driverID": id,
	})
	driver, err := s.repo.GetByIdWithEagerLoading(ctx, id)
	if err != nil {
		s.logger.Error("[DRIVER] GetByIdWithEagerLoading - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return driver, nil
	}

	return driver, nil
}

func (s *Service) GetByUserIdWithEagerLoading(ctx context.Context, userId int64) (*Driver, error) {
	s.logger.Debug("[DRIVER] GetByUserIdWithEagerLoading - DEBUG: ", map[string]any{
		"userID": userId,
	})
	driver, err := s.repo.GetByUserIdWithEagerLoading(ctx, userId)
	if err != nil {
		s.logger.Error("[DRIVER] GetByUserIdWithEagerLoading - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return driver, nil
	}

	return driver, nil
}

func (s *Service) List(ctx context.Context, specification *DriverSpecification) (*[]Driver, error) {
	s.logger.Debug("[DRIVER] List - DEBUG: ", map[string]any{
		"specification": specification,
	})
	drivers, err := s.repo.List(ctx, specification)
	if err != nil {
		s.logger.Error("[DRIVER] List - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return drivers, nil
}

func (s *Service) ListWithEagerLoading(ctx context.Context, specification *DriverSpecification) (*[]Driver, error) {
	s.logger.Debug("[DRIVER] ListWithEagerLoading - DEBUG: ", map[string]any{
		"specification": specification,
	})
	drivers, err := s.repo.ListWithEagerLoading(ctx, specification)
	if err != nil {
		s.logger.Error("[DRIVER] ListWithEagerLoading - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return drivers, nil
}

func (s *Service) Create(ctx context.Context, d *Driver) (int64, error) {
	s.logger.Debug("[DRIVER] Create - DEBUG: ", map[string]any{
		"driver": d,
	})
	driverID, err := s.repo.Create(ctx, d)
	if err != nil {
		s.logger.Error("[DRIVER] Create - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return 0, err
	}

	return driverID, nil
}

func (s *Service) Update(ctx context.Context, d *Driver) error {
	s.logger.Debug("[DRIVER] Update - DEBUG: ", map[string]any{
		"driver": d,
	})
	err := s.repo.Update(ctx, d)
	if err != nil {
		s.logger.Error("[DRIVER] Update - ERROR: ", map[string]any{
			"err": err.Error(),
		})

		return err
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	s.logger.Debug("[DRIVER] Delete - DEBUG: ", map[string]any{
		"driverID": id,
	})
	err := s.repo.Delete(ctx, id)
	if err != nil {
		s.logger.Error("[DRIVER] Delete - ERROR: ", map[string]any{
			"err": err.Error(),
		})

		return err
	}

	return nil
}
