package postgres

import (
	"log"

	"template-service/config"

	"github.com/rubenv/sql-migrate"
)

func RunMigration(cfg *config.Config) error {

	migration := migrate.FileMigrationSource{
		Dir: cfg.Postgres.MigrationDir,
	}

	conn := GetClient(cfg)

	migrate.SetSchema(cfg.Postgres.Schema)
	migrate.SetTable(cfg.Postgres.MigrationTable)

	_, err := migrate.Exec(conn.DB, cfg.Postgres.Dialect, migration, migrate.Up)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
