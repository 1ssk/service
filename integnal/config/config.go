package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/test/pkg/postgres"
)

type Config struct {
	Postgres postgres.Config `yaml:"POSTGRES" env:"POSTGRES" env-default:"POSTGRES"`

	GRPCPort int `yaml:"GRPC_PORT" env:"GRPC_PORT" env-default:"50051"`
}

func New() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig("./config/config.yml", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
