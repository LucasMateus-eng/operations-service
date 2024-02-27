package gin

import (
	"context"
	"log/slog"

	"github.com/LucasMateus-eng/operations-service/driver"
	postgres_driver "github.com/LucasMateus-eng/operations-service/driver/postgres"

	"github.com/LucasMateus-eng/operations-service/internal/db/postgres"
	"github.com/gin-gonic/gin"
)

func Handlers(ctx context.Context, logger *slog.Logger) *gin.Engine {
	db := postgres.InitPostgreSQL()
	driverRepo := postgres_driver.New(db)
	driverService := driver.NewService(driverRepo, logger)

	r := gin.Default()

	v1 := r.Group("v1")
	dGroup := v1.Group("drivers")
	{
		dGroup.GET("/", listDrivers(ctx, *driverService, logger))
		dGroup.POST("/", createDriver(ctx, *driverService, logger))
		dGroup.PUT("/:id", updateDriver(ctx, *driverService, logger))
		dGroup.DELETE("/:id", deleteDriver(ctx, *driverService, logger))
	}

	r.GET("/health", healthHandler)

	return r
}
