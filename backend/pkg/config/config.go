package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port        string `envconfig:"PET_SERVER_PORT" default:"4000"`
	DatabaseURI string `envconfig:"PET_DATABASE_URI" required:"true"`
}

var AppConfig Config

// LoadConfig loads the configuration from environment variables.
func LoadConfig() error {
	err := envconfig.Process("", &AppConfig)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return nil
}
