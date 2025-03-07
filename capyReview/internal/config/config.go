package config

import (
	"APIGateway/internal/models"
	"fmt"

	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Env      *EnvConfig
	Services *ServiceConfig
}

type EnvConfig struct {
	Port      string
	SecretKey string
}

type ServiceConfig struct {
	Services map[string]models.Service `yaml:"services"`
}

func New(yamlPath string) *Config {
	services, err := NewServiceConfig(yamlPath)
	if err != nil {
		log.Error().Msgf("%v", err)
	}
	return &Config{
		Env:      NewEnvConfig(),
		Services: services,
	}
}

func NewEnvConfig() *EnvConfig {
	return &EnvConfig{
		Port:      getStringEnv("port", ":8080"),
		SecretKey: getStringEnv("secret_key", ""),
	}
}

func NewServiceConfig(path string) (*ServiceConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read service config file: %w", err)
	}

	var config ServiceConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal service config file: %w", err)
	}

	return &config, nil
}

func getStringEnv(key string, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	} else {
		return defaultVal
	}
}
