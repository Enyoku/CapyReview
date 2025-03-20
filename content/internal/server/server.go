package server

import (
	"contentService/internal/api"
	"contentService/internal/config"

	"github.com/rs/zerolog/log"
)

type Server struct {
	port   string
	api    *api.API
	config *config.Config
}

func New() (*Server, error) {
	// Инициализация структуры конфигурации
	config, err := config.New()
	if err != nil {
		log.Fatal().Msg("")
	}

	// Инициализация API
	api, err := api.New()
	if err != nil {
		log.Fatal().Msg("")
	}

	return &Server{
		port:   "",
		api:    api,
		config: config,
	}, nil
}

func (s *Server) Run() {
	s.api.Run(s.config.Port)
}
