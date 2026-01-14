package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	App      App
	Context  Context
	Security Security
	Postgres Postgres
	GRPC     GRPC
}

type App struct {
	Name    string `env:"APP_NAME,required"`
	Version string `env:"APP_VERSION,required"`
}

type Context struct {
	Timeout time.Duration `env:"APP_CONTEXT_TIMEOUT" envDefault:"5s"`
}

type Security struct {
	BcryptCost int `env:"SECURITY_BCRYPT_COST" envDefault:"10"`
}

type Postgres struct {
	URL     string `env:"POSTGRES_URL,required"`
	PoolMax int    `env:"POSTGRES_POOL_MAX,required"`
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
