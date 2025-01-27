package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Token         string `yaml:"token"`
	PoolerTimeout int    `yaml:"pooler_timeout"`
	PgConnectURL  string `yaml:"pg_connect_string"`
	LogLevel      string `yaml:"log_level"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadConfig("/config.yaml", cfg); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	return cfg, nil
}
