package main

import (
	"context"
	"log"

	"github.com/LucasMateus-eng/operations-service/config"
	driver_vehicle_postgres_dto "github.com/LucasMateus-eng/operations-service/driver-vehicle/postgres/dto"
	"github.com/LucasMateus-eng/operations-service/internal/api"
	"github.com/LucasMateus-eng/operations-service/internal/db/postgres"
	"github.com/LucasMateus-eng/operations-service/internal/http/gin"
	"github.com/LucasMateus-eng/operations-service/internal/logging"
)

func main() {
	db := postgres.InitPostgreSQL()
	db.RegisterModel((*driver_vehicle_postgres_dto.DriverVehicleDTO)(nil))

	ctx := context.Background()
	config := config.NewConfig()
	logger := logging.InitializerLogging(config)

	h := gin.Handlers(ctx, db, logger)
	err := api.Start(config.AppDefaultPort, logger, h)
	if err != nil {
		log.Fatal("error when initializing an application")
	}
}
