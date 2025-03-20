package db

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Создаёт новый клиент MongoDB
func NewMongoClient(uri string) (*mongo.Client, error) {
	// Устанавливаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Настройка клиента MongoDB
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Error().Msgf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}

	// Проверяем подключение к базе данных
	if err := client.Ping(ctx, nil); err != nil {
		log.Error().Msgf("Failed to ping MongoDB: %v", err)
		return nil, err
	}

	log.Info().Msgf("Successfully connected to MongoDB")
	return client, nil
}
