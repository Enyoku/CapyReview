package server

import (
	"contentService/internal/api"
	"contentService/internal/config"
	"contentService/internal/db"

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

	//
	db, err := db.NewMongoClient(config.MongoURI)
	if err != nil {
		return nil, err
	}

	// Инициализация зависимостей
	movieService, err := initializeDependencies(db)
	if err != nil {
		return nil, err
	}

	// Создание API
	api, err := api.New(movieService)
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
