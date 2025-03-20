package main

import (
	"contentService/internal/server"

	"github.com/rs/zerolog/log"
)

func main() {
	// Инициализация сервера
	s, err := server.New()
	if err != nil {
		log.Fatal().Msg("")
	}

	// Запуск сервера
	s.Run()
}
