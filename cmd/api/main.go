package main

import (
	"context"
	"log"

	"github.com/LucasMateus-eng/operations-service/config"
	"github.com/LucasMateus-eng/operations-service/internal/api"
	"github.com/LucasMateus-eng/operations-service/internal/db/postgres"
	"github.com/LucasMateus-eng/operations-service/internal/http/gin"
	"github.com/LucasMateus-eng/operations-service/internal/logging"
)

func main() {
	ctx := context.Background()
	config := config.NewConfig()
	db := postgres.InitPostgreSQL()
	logger := logging.InitializerLogging(config)

	h := gin.Handlers(ctx, db, logger)
	err := api.Start(config.AppDefaultPort, logger, h)
	if err != nil {
		log.Fatal("error when initializing an application")
	}
}
