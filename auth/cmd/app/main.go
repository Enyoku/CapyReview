package main

import (
	"authService/internal/server"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("no .env files")
	}
}

func main() {
	s, err := server.New()
	if err != nil {
		log.Fatal().Msg("Server error")
	}
	s.Run()
}
