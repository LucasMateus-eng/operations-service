package gin

import (
	"context"

	"github.com/LucasMateus-eng/operations-service/driver"
	drivervehicle "github.com/LucasMateus-eng/operations-service/driver-vehicle"
	postgres_driver_vehicle "github.com/LucasMateus-eng/operations-service/driver-vehicle/postgres"
	postgres_driver "github.com/LucasMateus-eng/operations-service/driver/postgres"
	"github.com/LucasMateus-eng/operations-service/internal/logging"
	"github.com/LucasMateus-eng/operations-service/user"
	postgres_user "github.com/LucasMateus-eng/operations-service/user/postgres"
	"github.com/LucasMateus-eng/operations-service/vehicle"
	postgres_vehicle "github.com/LucasMateus-eng/operations-service/vehicle/postgres"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func Handlers(ctx context.Context, db *bun.DB, logger *logging.Logging) *gin.Engine {
	userRepo := postgres_user.New(db)
	userService := user.NewService(userRepo, logger)
	driverRepo := postgres_driver.New(db)
	driverService := driver.NewService(driverRepo, logger)
	vehicleRepo := postgres_vehicle.New(db)
	vehicleService := vehicle.NewService(vehicleRepo, logger)
	driverVehicleRepo := postgres_driver_vehicle.New(db)
	driverVehicleService := drivervehicle.NewService(driverVehicleRepo, logger)

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
		dGroup.GET("/:id", getDriver(ctx, driverService, logger))
		dGroup.PUT("/:id", updateDriver(ctx, driverService, logger))
		dGroup.DELETE("/:id", deleteDriver(ctx, driverService, logger))
	}

	vGroup := v1.Group("vehicles")
	{
		vGroup.GET("/", listVehicles(ctx, vehicleService, logger))
		vGroup.POST("/", createVehicle(ctx, vehicleService, logger))
		vGroup.GET("/:id", getVehicle(ctx, vehicleService, logger))
		vGroup.PUT("/:id", updateVehicle(ctx, vehicleService, logger))
		vGroup.DELETE("/:id", deleteVehicle(ctx, vehicleService, logger))
	}

	dvGroup := v1.Group("drivers-vehicles")
	{
		dvGroup.POST("/", createDriverVehicle(ctx, driverVehicleService, logger))
		dvGroup.GET("/:driver_id", listVehiclesByDriverID(ctx, driverVehicleService, logger))
		dvGroup.GET("/:vehicle_id", listDriversByVehicleID(ctx, driverVehicleService, logger))
		dvGroup.DELETE("/:driver_id/:vehicle_id", deleteDriverVehicle(ctx, driverVehicleService, logger))
	}

	r.GET("/health", healthHandler)

	return r
}
