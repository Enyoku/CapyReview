package config

import "os"

type Config struct {
	Port string
	*DBConfig
}

type DBConfig struct {
	Port     string
	User     string
	Password string
	DbName   string
}

func New() (*Config, error) {
	return &Config{
		Port: getStringEnv("port", ":9002"),
		DBConfig: &DBConfig{
			Port:     getStringEnv("port", "27017"),
			User:     getStringEnv("port", ""),
			Password: getStringEnv("port", ""),
			DbName:   getStringEnv("port", "contentService"),
		},
	}, nil
}

func getStringEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	} else {
		return defaultVal
	}
}
