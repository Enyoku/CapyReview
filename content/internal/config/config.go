package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port string
	*DBConfig
	MongoURI string
}

type DBConfig struct {
	Port     string
	User     string
	Password string
	DbName   string
	Host     string
}

func New() (*Config, error) {
	// Получаем значения из переменных окружения
	dbConfig := &DBConfig{
		Port:     getStringEnv("MONGO_PORT", "27017"),
		User:     getStringEnv("MONGO_USER", ""),
		Password: getStringEnv("MONGO_PASSWORD", ""),
		DbName:   getStringEnv("MONGO_DB_NAME", "contentService"),
		Host:     getStringEnv("mongo_host", "localhost"),
	}

	// Формируем MongoURI с использованием значений из DBConfig
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
	)

	// Если пользователь и пароль не указаны, убираем аутентификацию из URI
	if dbConfig.User == "" || dbConfig.Password == "" {
		mongoURI = fmt.Sprintf("mongodb://%s:%s/%s",
			getStringEnv("MONGO_HOST", "localhost"),
			dbConfig.Port,
			dbConfig.DbName,
		)
	}

	return &Config{
		Port:     getStringEnv("PORT", ":9002"),
		DBConfig: dbConfig,
		MongoURI: mongoURI,
	}, nil
}

func getStringEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	} else {
		return defaultVal
	}
}
