package config

import (
	"os"
	"time"
)

type Config struct {
	Port string
	DB   *DBConfig
	JWT  *JWT
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

type JWT struct {
	JWTSecretKey     string
	JWTExpiresIn     time.Duration
	RefreshExpiresIn time.Duration
}

func New() (*Config, error) {
	day := 24 * time.Hour
	return &Config{
		Port: getStringEnv("port", ":9001"),
		DB: &DBConfig{
			Username: getStringEnv("postgres_user", ""),
			Password: getStringEnv("postgres_password", ""),
			Host:     getStringEnv("postgres_host", ""),
			Port:     getStringEnv("postgres_port", ""),
			Name:     getStringEnv("postgres_name", ""),
		},
		JWT: &JWT{
			JWTSecretKey:     getStringEnv("jwt_secret_key", ""),
			JWTExpiresIn:     day,
			RefreshExpiresIn: 7 * day,
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
