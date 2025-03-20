package server

import (
	"contentService/internal/api"

	"github.com/rs/zerolog/log"
)

type Server struct {
	port string
	api  *api.API
}

func New() (*Server, error) {

	// Инициализация API
	api, err := api.New()
	if err != nil {
		log.Fatal().Msg("")
	}

	return &Server{
		port: "",
		api:  api,
	}, nil
}

func (s *Server) Run() {
	s.api.Run("")
}
