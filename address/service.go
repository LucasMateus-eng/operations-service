package address

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

func (s *Service) GetById(ctx context.Context, id int64) (*Address, error) {
	address, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (s *Service) GetByUserID(ctx context.Context, userID int64) (*Address, error) {
	address, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (s *Service) Create(ctx context.Context, a *Address) (int64, error) {
	addressID, err := s.repo.Create(ctx, a)
	if err != nil {
		return 0, err
	}

	return addressID, nil
}

func (s *Service) Update(ctx context.Context, a *Address) error {
	return s.repo.Update(ctx, a)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
