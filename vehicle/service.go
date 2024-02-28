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

func (s *Service) GetByVehicleId(ctx context.Context, id int64) (*Vehicle, error) {
	vehicle, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (s *Service) GetByPlate(ctx context.Context, plate string) (*Vehicle, error) {
	vehicle, err := s.repo.GetByPlate(ctx, plate)
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (s *Service) GetByRenavam(ctx context.Context, renavam string) (*Vehicle, error) {
	vehicle, err := s.repo.GetByRenavam(ctx, renavam)
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (s *Service) List(ctx context.Context, specification *VehicleSpectification) (*[]Vehicle, error) {
	vehicles, err := s.repo.List(ctx, specification)
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (s *Service) Create(ctx context.Context, v *Vehicle) (int64, error) {
	vehicleID, err := s.repo.Create(ctx, v)
	if err != nil {
		return 0, err
	}

	return vehicleID, nil
}

func (s *Service) Update(ctx context.Context, v *Vehicle) error {
	return s.repo.Update(ctx, v)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
