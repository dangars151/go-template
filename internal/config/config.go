package config

import "fmt"

type Config struct {
	Postgres *PostgresConfig
	Port     uint32
	Host     string
}

type PostgresConfig struct {
	Host     string
	Port     uint32
	User     string
	Password string
	Database string
}

func (pc *PostgresConfig) GetPGConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		pc.User, pc.Password, pc.Host, pc.Port, pc.Database)
}
