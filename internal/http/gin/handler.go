package gin

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func Handlers(ctx context.Context, logger *slog.Logger) *gin.Engine {
	r := gin.Default()
	r.GET("/health", healthHandler)

	return r
}
