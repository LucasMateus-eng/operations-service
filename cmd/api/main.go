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

const (
	DEFAULT_CONFIG_TYPE = "env"
	DEFAULT_CONFIG_FILE = ".env"
	DEFAULT_CONFIG_PATH = "./"
)

func main() {
	ctx := context.Background()
	config := config.NewConfig(DEFAULT_CONFIG_TYPE, DEFAULT_CONFIG_FILE, DEFAULT_CONFIG_PATH)

	db := postgres.InitPostgreSQL(config)
	logger := logging.InitializerLogging(config)

	h := gin.Handlers(ctx, db, logger)
	err := api.Start(config.AppDefaultPort, logger, h)
	if err != nil {
		log.Fatalf("error when initializing an application: %s", err.Error())
	}
}
