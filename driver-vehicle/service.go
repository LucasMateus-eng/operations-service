package drivervehicle

import "log/slog"

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
