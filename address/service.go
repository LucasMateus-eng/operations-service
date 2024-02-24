package address

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

func (s *Service) GetById(ctx context.Context, id int) (*Address, error) {
	address, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (s *Service) GetByUserID(ctx context.Context, userID int) (*Address, error) {
	address, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (s *Service) Create(ctx context.Context, a *Address) (int, error) {
	addressID, err := s.repo.Create(ctx, a)
	if err != nil {
		return 0, err
	}

	return addressID, nil
}

func (s *Service) Update(ctx context.Context, a *Address) error {
	return s.repo.Update(ctx, a)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
