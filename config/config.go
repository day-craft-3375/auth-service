package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	App  App
	GRPC GRPC
}

type App struct {
	Name    string `env:"APP_NAME,required"`
	Version string `env:"APP_VERSION,required"`
}

type GRPC struct {
	Port string `env:"GRPC_PORT,required"`
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения .env файла: %w", err)
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("ошибка конфигурации: %w", err)
	}

	return cfg, nil
}
