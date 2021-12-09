package config

import "github.com/kelseyhightower/envconfig"

type db struct {
	Host     string `envconfig:"APP_DB_HOST" default:"localhost"`
	Port     string `envconfig:"APP_DB_PORT" default:"5432"`
	User     string `envconfig:"APP_DB_USER" default:"db"`
	Password string `envconfig:"APP_DB_PASSWORD" default:"db"`
	Database string `envconfig:"APP_DB_DATABASE" default:"banking"`
}

type Config struct {
	DB db
}

func Environ() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	return cfg, err
}
