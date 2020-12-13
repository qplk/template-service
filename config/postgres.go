package config

type Postgres struct {
	Host           string `env:"PG_HOST"`
	Port           int    `env:"PG_PORT"`
	User           string `env:"PG_USER"`
	Pass           string `env:"PG_PASS"`
	Database       string `env:"PG_DB"`
	Schema         string `env:"PG_SCHEMA"`
	Dialect        string `env:"PG_DIALECT"`
	MigrationDir   string `env:"PG_MIGRATION_DIR"`
	MigrationTable string `env:"PG_MIGRATION_TABLE"`
}
