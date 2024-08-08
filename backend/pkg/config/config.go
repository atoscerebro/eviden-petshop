package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServicePort  string `envconfig:"PET_SERVER_PORT" default:"4000"`
	DatabaseName string `envconfig:"PET_DB_NAME" required:"true"`
	DatabaseUser string `envconfig:"PET_DB_USER" required:"true"`
	DatabasePass string `envconfig:"PET_DB_PASS" required:"true"`
	DatabaseHost string `envconfig:"PET_DB_HOST" required:"true"`
	DatabasePort string `envconfig:"PET_DB_PORT" required:"true"`
	DatabaseURI  string
}

var AppConfig Config

// LoadConfig loads the environment variables into the Config struct
func LoadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	// encode the database credentials to avoid special characters
	cfg.DatabaseURI = fmt.Sprintf("postgresql://%s:%s@%s:%s",
		cfg.DatabaseUser, cfg.DatabasePass, cfg.DatabaseHost, cfg.DatabasePort)

	return &cfg, nil
}
