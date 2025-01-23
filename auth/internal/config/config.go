package config

import "os"

type Config struct {
	Port string
	DB   *DBConfig
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Name     string
}

func New() (*Config, error) {
	return &Config{
		Port: getStringEnv("port", ":9001"),
		DB: &DBConfig{
			Username: getStringEnv("postgres_user", ""),
			Password: getStringEnv("postgres_password", ""),
			Host:     getStringEnv("postgres_host", ""),
			Name:     getStringEnv("postgres_name", ""),
		},
	}, nil
}

func getStringEnv(key string, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	} else {
		return defaultVal
	}
}
