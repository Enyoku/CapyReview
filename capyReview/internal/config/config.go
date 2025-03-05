package config

import (
	"APIGateway/internal/models"
	"os"
)

type Config struct {
	Port      string
	SecretKey string
}

type ServiceConfig struct {
	Services map[string]models.Service `yaml:"services"`
}

func New() *Config {
	return &Config{
		Port:      getStringEnv("port", ":8080"),
		SecretKey: getStringEnv("secret_key", ""),
	}
}

func getStringEnv(key string, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	} else {
		return defaultVal
	}
}
