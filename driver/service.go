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
	driver, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (s *Service) GetByUserId(ctx context.Context, userId int64) (*Driver, error) {
	driver, err := s.repo.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (s *Service) GetByIdWithEagerLoading(ctx context.Context, id int64) (*Driver, error) {
	driver, err := s.repo.GetByIdWithEagerLoading(ctx, id)
	if err != nil {
		return driver, nil
	}

	return driver, nil
}

func (s *Service) GetByUserIdWithEagerLoading(ctx context.Context, userId int64) (*Driver, error) {
	driver, err := s.repo.GetByUserIdWithEagerLoading(ctx, userId)
	if err != nil {
		return driver, nil
	}

	return driver, nil
}

func (s *Service) List(ctx context.Context, specification *DriverSpecification) (*[]Driver, error) {
	drivers, err := s.repo.List(ctx, specification)
	if err != nil {
		return nil, err
	}

	return drivers, nil
}

func (s *Service) ListWithEagerLoading(ctx context.Context, specification *DriverSpecification) (*[]Driver, error) {
	drivers, err := s.repo.ListWithEagerLoading(ctx, specification)
	if err != nil {
		return nil, err
	}

	return drivers, nil
}

func (s *Service) Create(ctx context.Context, d *Driver) (int64, error) {
	driverID, err := s.repo.Create(ctx, d)
	if err != nil {
		return 0, err
	}

	return driverID, nil
}

func (s *Service) Update(ctx context.Context, d *Driver) error {
	return s.repo.Update(ctx, d)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
