package main

import (
	"APIGateway/internal/server"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("no .env files")
	}
}

func main() {
	s := server.New()
	s.Run()
}
