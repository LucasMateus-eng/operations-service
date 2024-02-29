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

func (s *Service) GetByID(ctx context.Context, id int64) (*Driver, error) {
	s.logger.Debug("[DRIVER] GetByID - DEBUG: ", map[string]any{
		"driverID": id,
	})
	driver, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("[DRIVER] GetByID - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return driver, nil
}

func (s *Service) GetByUserID(ctx context.Context, userId int64) (*Driver, error) {
	s.logger.Debug("[DRIVER] GetByUserID - DEBUG: ", map[string]any{
		"userID": userId,
	})
	driver, err := s.repo.GetByUserID(ctx, userId)
	if err != nil {
		s.logger.Error("[DRIVER] GetByUserID - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return driver, nil
}

func (s *Service) GetByIDWithEagerLoading(ctx context.Context, id int64) (*Driver, error) {
	s.logger.Debug("[DRIVER] GetByIDWithEagerLoading - DEBUG: ", map[string]any{
		"driverID": id,
	})
	driver, err := s.repo.GetByIDWithEagerLoading(ctx, id)
	if err != nil {
		s.logger.Error("[DRIVER] GetByIDWithEagerLoading - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return driver, nil
}

func (s *Service) GetByUserIDWithEagerLoading(ctx context.Context, userId int64) (*Driver, error) {
	s.logger.Debug("[DRIVER] GetByUserIDWithEagerLoading - DEBUG: ", map[string]any{
		"userID": userId,
	})
	driver, err := s.repo.GetByUserIDWithEagerLoading(ctx, userId)
	if err != nil {
		s.logger.Error("[DRIVER] GetByUserIDWithEagerLoading - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
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
