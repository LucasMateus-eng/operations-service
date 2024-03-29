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

func (s *Service) GetByID(ctx context.Context, id int64) (*User, error) {
	s.logger.Debug("[USER] GetByID - DEBUG: ", map[string]any{
		"userID": id,
	})
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("[USER] GetByID - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return user, nil
}

func (s *Service) GetByUsername(ctx context.Context, username string) (*User, error) {
	s.logger.Debug("[USER] GetByUsername - DEBUG: ", map[string]any{
		"userUsername": username,
	})
	user, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		s.logger.Error("[USER] GetByUsername - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return user, nil
}

func (s *Service) GetByRole(ctx context.Context, role Role) (*User, error) {
	s.logger.Debug("[USER] GetByRole - DEBUG: ", map[string]any{
		"userRole": role,
	})
	user, err := s.repo.GetByRole(ctx, role)
	if err != nil {
		s.logger.Error("[USER] GetByRole - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return user, nil
}

func (s *Service) Create(ctx context.Context, u *User) (int64, error) {
	s.logger.Debug("[USER] Create - DEBUG: ", map[string]any{
		"user": u,
	})
	userID, err := s.repo.Create(ctx, u)
	if err != nil {
		s.logger.Error("[USER] Create - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return 0, err
	}

	return userID, nil
}

func (s *Service) Update(ctx context.Context, u *User) error {
	s.logger.Debug("[USER] Update - DEBUG: ", map[string]any{
		"user": u,
	})
	err := s.repo.Update(ctx, u)
	if err != nil {
		s.logger.Error("[USER] Update - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return err
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	s.logger.Debug("[USER] Delete - DEBUG: ", map[string]any{
		"userID": id,
	})
	err := s.repo.Delete(ctx, id)
	if err != nil {
		s.logger.Error("[USER] Delete - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return err
	}

	return nil
}
