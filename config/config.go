package config

import (
	"github.com/jinzhu/configor"
)

type HttpServerConfig struct {
	Address string `default:0.0.0.0"`
	Port    int
}

type Log struct {
	Level string `default:"info"`
}

type Http struct {
	Server HttpServerConfig
}

type UserRepository struct {
	Implementation string
	Config         string
}

type Config struct {
	Log  Log
	Http Http
	UserRepository
}

func New() (*Config, error) {
	var cfg Config
	err := configor.New(
		&configor.Config{
			ENVPrefix: "APP",
			Verbose:   false,
			Debug:     false,
		},
	).Load(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, err
}
