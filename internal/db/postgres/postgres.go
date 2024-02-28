package postgres

import (
	"database/sql"
	"fmt"

	"github.com/LucasMateus-eng/operations-service/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func buildPostgreSQLConnDSN(config *config.Config) string {
	username := config.DBUser
	password := config.DBPass
	host := config.DBHost
	port := config.DBPort
	database := config.DBName

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&timezone=America/Sao_Paulo", username, password, host, port, database)

	return dsn
}

func InitPostgreSQL(config *config.Config) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(buildPostgreSQLConnDSN(config))))

	db := bun.NewDB(sqldb, pgdialect.New())

	return db
}
