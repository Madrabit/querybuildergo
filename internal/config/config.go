package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB     DBConfig
	Server ServerConfig
}

type DBConfig struct {
	Server   string `envconfig:"SERVER" required:"true"`
	Port     int    `envconfig:"PORT" required:"true"`
	User     string `envconfig:"USER" required:"true"`
	Pass     string `envconfig:"PASS" required:"true"`
	Database string `envconfig:"DATABASE" required:"true"`
}

type ServerConfig struct {
	Address string `envconfig:"ADDRESS" required:"true"`
	Port    int    `envconfig:"PORT" required:"true"`
}

func Load() (*Config, error) {
	var cfg Config
	if db, err := LoadDbConfig(); err != nil {
		return nil, err
	} else {
		cfg.DB = *db
	}
	if server, err := LoadServerConfig(); err != nil {
		return nil, err
	} else {
		cfg.Server = *server
	}
	return &cfg, nil
}

func LoadDbConfig() (*DBConfig, error) {
	var cfg DBConfig
	err := envconfig.Process("DB", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func LoadServerConfig() (*ServerConfig, error) {
	var cfg ServerConfig
	err := envconfig.Process("SERVER", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
