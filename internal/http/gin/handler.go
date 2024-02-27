package gin

import (
	"context"
	"log/slog"

	"github.com/LucasMateus-eng/operations-service/driver"
	postgres_driver "github.com/LucasMateus-eng/operations-service/driver/postgres"
	"github.com/LucasMateus-eng/operations-service/internal/db/postgres"
	"github.com/LucasMateus-eng/operations-service/user"
	postgres_user "github.com/LucasMateus-eng/operations-service/user/postgres"
	"github.com/gin-gonic/gin"
)

func Handlers(ctx context.Context, logger *slog.Logger) *gin.Engine {
	db := postgres.InitPostgreSQL()

	userRepo := postgres_user.New(db)
	userService := user.NewService(userRepo, logger)
	driverRepo := postgres_driver.New(db)
	driverService := driver.NewService(driverRepo, logger)

	r := gin.Default()

	v1 := r.Group("v1")
	uGroup := v1.Group("/users")
	{
		uGroup.POST("/", createUser(ctx, userService, logger))
		uGroup.GET(":id", getUser(ctx, userService, logger))
		uGroup.PUT(":id", updateUser(ctx, userService, logger))
		uGroup.DELETE(":id", deleteUser(ctx, userService, logger))
	}

	dGroup := v1.Group("drivers")
	{
		dGroup.GET("/", listDrivers(ctx, driverService, logger))
		dGroup.POST("/", createDriver(ctx, driverService, logger))
		dGroup.PUT("/:id", updateDriver(ctx, driverService, logger))
		dGroup.DELETE("/:id", deleteDriver(ctx, driverService, logger))
	}

	r.GET("/health", healthHandler)

	return r
}
