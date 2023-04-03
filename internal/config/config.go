package config

import "github.com/caarlos0/env/v7"

type Config struct {
	AppMode string `env:"APP_MODE" envDefault:"debug"`
	AppPort string `env:"APP_PORT" envDefault:"8080"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
