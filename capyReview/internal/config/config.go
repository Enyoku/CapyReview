package config

import "APIGateway/internal/models"

type Config struct {
	Host      string
	Port      string
	SecretKey string
}

type ServiceConfig struct {
	Services map[string]models.Service `yaml:"services"`
}

func New() *Config {
	return &Config{
		Port: ":8080",
	}
}
