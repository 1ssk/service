package config

import (
	"github.com/1ssk/service/pkg/postgres"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres postgres.Config `yaml:"POSTGRES" env:"POSTGRES" env-default:"POSTGRES"`

	GRPCPort int `yaml:"GRPC_PORT" env:"GRPC_PORT" env-default:"50051"`
	RestPort int `yaml:"REST_PORT" env:"REST_PORT" env-default:"8081"`
}

func New() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig("./config/config.yml", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
