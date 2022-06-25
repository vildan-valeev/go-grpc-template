package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
	"go-grpc-template/pkg/database"
)

type Config struct {
	LogLevel string `env:"LOGLEVEL,required" envDefault:"debug"` // debug, info, warn, error, fatal, ""
	IP       string `env:"IP,required" envDefault:"0.0.0.0"`
	GRPCPort string `env:"HTTP_PORT,required" envDefault:"8000"`
	BaseURL  string `env:"BASE_URL,required" envDefault:"http://localhost"`

	Database database.Config
}

// New создает экземпляр Config и заполняет его значениями переменных окружения.
func New() *Config {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		log.Fatal().Err(err).Msg("parse env")
	}

	return cfg
}
