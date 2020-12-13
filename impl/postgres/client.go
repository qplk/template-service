package postgres

import (
	"fmt"
	"log"

	"template-service/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var client *sqlx.DB

func GetClient(cfg *config.Config) *sqlx.DB {
	if !isClientActive() {
		conn, err := sqlx.Open(cfg.Postgres.Dialect, getDataSource(cfg))
		if err != nil {
			log.Fatal(err)
		}

		client = conn
	}

	return client
}

func isClientActive() bool {
	if client == nil {
		return false
	}

	if client.Ping() != nil {
		return false
	}

	return true
}

func getDataSource(cfg *config.Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable binary_parameters=yes",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Pass, cfg.Postgres.Database)
}
