package user

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

func (s *Service) GetById(ctx context.Context, id int) (*User, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetByUsername(ctx context.Context, username string) (*User, error) {
	user, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetByRole(ctx context.Context, role Role) (*User, error) {
	user, err := s.repo.GetByRole(ctx, role)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Create(ctx context.Context, u *User) (int, error) {
	userID, err := s.repo.Create(ctx, u)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (s *Service) Update(ctx context.Context, u *User) error {
	return s.repo.Update(ctx, u)

}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
