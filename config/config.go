package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type (
	Config struct {
		Postgres
		GRPC
		Gateway
	}

	Postgres struct {
		URL string `env:"POSTGRES_URL"`
	}

	GRPC struct {
		Host string `env:"GRPC_HOST"`
		Port string `env:"GRPC_PORT"`
	}

	Gateway struct {
		Port string `env:"GATEWAY_PORT"`
		Host string `env:"GATEWAY_HOST"`
	}
)

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.Wrap(err, "load env file")
	}

	cfg := &Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, errors.Wrap(err, "parse env file")
	}

	return cfg, nil
}
