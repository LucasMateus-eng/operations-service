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

func (s *Service) GetByID(ctx context.Context, id int64) (*Address, error) {
	s.logger.Debug("[ADDRESS] GetByID - DEBUG: ", map[string]any{
		"addressID": id,
	})
	address, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("[ADDRESS] GetByID - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return address, nil
}

func (s *Service) GetByUserID(ctx context.Context, userID int64) (*Address, error) {
	s.logger.Debug("[ADDRESS] GetByUserID - DEBUG: ", map[string]any{
		"userID": userID,
	})
	address, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		s.logger.Error("[ADDRESS] GetByUserID - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return nil, err
	}

	return address, nil
}

func (s *Service) Create(ctx context.Context, a *Address) (int64, error) {
	s.logger.Debug("[ADDRESS] Create - DEBUG: ", map[string]any{
		"address": a,
	})
	addressID, err := s.repo.Create(ctx, a)
	if err != nil {
		s.logger.Error("[ADDRESS] Create - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return 0, err
	}

	return addressID, nil
}

func (s *Service) Update(ctx context.Context, a *Address) error {
	s.logger.Debug("[ADDRESS] Update - DEBUG: ", map[string]any{
		"address": a,
	})
	err := s.repo.Update(ctx, a)
	if err != nil {
		s.logger.Error("[ADDRESS] Update - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return err
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	s.logger.Debug("[ADDRESS] Delete - DEBUG: ", map[string]any{
		"addressID": id,
	})
	err := s.repo.Delete(ctx, id)
	if err != nil {
		s.logger.Error("[ADDRESS] Delete - ERROR: ", map[string]any{
			"err": err.Error(),
		})
		return err
	}

	return nil
}
