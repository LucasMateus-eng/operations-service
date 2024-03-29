package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LucasMateus-eng/operations-service/internal/logging"
)

const TIMEOUT = 30 * time.Second

type ServerOption func(server *http.Server)

// Start a new http server with graceful shutdown and default parameters
func Start(port string, logger *logging.Logging, handler http.Handler, options ...ServerOption) error {

	srv := &http.Server{
		ReadTimeout:  TIMEOUT,
		WriteTimeout: TIMEOUT,
		Addr:         ":" + port,
		Handler:      handler,
	}

	for _, o := range options {
		o(srv)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func() {
		<-ctx.Done()
		logger.Info("stopping server", nil)
		err := srv.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}()

	logger.Info("server started successfully", map[string]any{
		"port": port,
	})
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

// WithReadTimeout configure http.Server parameter ReadTimeout
func WithReadTimeout(t time.Duration) ServerOption {
	return func(srv *http.Server) {
		srv.ReadTimeout = t
	}
}

// WithWriteTimeout configure http.Server parameter WriteTimeout
func WithWriteTimeout(t time.Duration) ServerOption {
	return func(srv *http.Server) {
		srv.WriteTimeout = t
	}
}
