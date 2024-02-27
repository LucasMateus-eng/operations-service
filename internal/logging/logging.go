package logging

import (
	"log/slog"
	"os"
	"strings"

	"github.com/LucasMateus-eng/operations-service/config"
)

type Logging struct {
	config *config.Config
	logger *slog.Logger
}

func InitializerLogging(
	config *config.Config,
) *Logging {
	opts := &slog.HandlerOptions{Level: parseLevel(config.AppLogLevel)}
	jsonHandler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(jsonHandler).With(
		"application_metadata",
		map[string]any{
			"name":        config.AppName,
			"environment": config.AppEnv,
		},
	)
	return &Logging{
		config,
		logger,
	}
}

func (l *Logging) Info(message string, data any) {
	l.logger.Info(message, slog.Any("data", data))
}

func (l *Logging) Warn(message string, data any) {
	l.logger.Warn(message, slog.Any("data", data))
}

func (l *Logging) Error(message string, data any) {
	l.logger.Error(message, slog.Any("data", data))
}

func (l *Logging) Debug(message string, data any) {
	l.logger.Debug(message, slog.Any("data", data))
}

func parseLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "error":
		return slog.LevelError
	case "warn":
		return slog.LevelWarn
	case "debug":
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}
