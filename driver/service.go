package driver

import (
	"context"
	"log/slog"
)

type Service struct {
	repo   Repository
	logger *slog.Logger
}

func NewService(r Repository, l *slog.Logger) *Service {
	return &Service{
		repo:   r,
		logger: l,
	}
}

func (s *Service) GetByDriverId(ctx context.Context, id int) (*Driver, error) {
	driver, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (s *Service) GetByUserId(ctx context.Context, userId int) (*Driver, error) {
	driver, err := s.repo.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (s *Service) GetByIdWithEagerLoading(ctx context.Context, id int) (*Driver, error) {
	driver, err := s.repo.GetByIdWithEagerLoading(ctx, id)
	if err != nil {
		return driver, nil
	}

	return driver, nil
}

func (s *Service) GetByUserIdWithEagerLoading(ctx context.Context, userId int) (*Driver, error) {
	driver, err := s.repo.GetByUserIdWithEagerLoading(ctx, userId)
	if err != nil {
		return driver, nil
	}

	return driver, nil
}

func (s *Service) List(ctx context.Context, specification DriverSpectification) (*[]Driver, error) {
	drivers, err := s.repo.List(ctx, specification)
	if err != nil {
		return nil, err
	}

	return drivers, nil
}

func (s *Service) ListWithEagerLoading(ctx context.Context, specification DriverSpectification) (*[]Driver, error) {
	drivers, err := s.repo.ListWithEagerLoading(ctx, specification)
	if err != nil {
		return nil, err
	}

	return drivers, nil
}

func (s *Service) Create(ctx context.Context, d *Driver) (int, error) {
	driverID, err := s.repo.Create(ctx, d)
	if err != nil {
		return 0, err
	}

	return driverID, nil
}

func (s *Service) Update(ctx context.Context, d *Driver) error {
	return s.repo.Update(ctx, d)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
