package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Auth     AuthConfig
}

type ServerConfig struct {
	Port         string        `envconfig:"PORT" default:"3000"`
	ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT" default:"30s"`
	WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"30s"`
	IdleTimeout  time.Duration `envconfig:"IDLE_TIMEOUT" default:"60s"`
}

type DatabaseConfig struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	User     string `envconfig:"DB_USER" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD" default:"password"`
	Name     string `envconfig:"DB_NAME" default:"taskdb"`
	SSLMode  string `envconfig:"DB_SSLMODE" default:"disable"`
}

type AuthConfig struct {
	Username string `envconfig:"BASIC_AUTH_USERNAME" default:"admin"`
	Password string `envconfig:"BASIC_AUTH_PASSWORD" default:"password123"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
