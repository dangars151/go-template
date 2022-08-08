package config

import "github.com/spf13/viper"

func LoadFromEnv() *Config {
	viper.AutomaticEnv()
	return &Config{
		Postgres: &PostgresConfig{
			Host:     viper.GetString("POSTGRESQL_HOST"),
			Port:     viper.GetUint32("POSTGRESQL_PORT"),
			User:     viper.GetString("POSTGRESQL_USERNAME"),
			Password: viper.GetString("POSTGRESQL_PASSWORD"),
			Database: viper.GetString("POSTGRESQL_DATABASE"),
		},
		Host: viper.GetString("HOST"),
		Port: viper.GetUint32("PORT"),
	}
}
