package main

import (
	"contentService/internal/server"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// Инициализация .env файла
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("no .env files")
	}
}

func main() {
	// Инициализация сервера
	s, err := server.New()
	if err != nil {
		log.Fatal().Msg("")
	}

	// Запуск сервера
	s.Run()
}
