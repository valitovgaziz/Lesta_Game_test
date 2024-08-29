package models

type PSQLConfig struct {
	PGUSER     string `env:"PGUSER"`
	PGPASSWORD string `env:"PGPASSWORD`
	PGDATABASE string `env:"PGDATABASE"`
	PGPORT     string `env:"PGPORT"`
	PGHOST     string `env:"PGHOST"`
}
