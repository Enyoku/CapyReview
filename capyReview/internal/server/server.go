package server

import (
	"APIGateway/internal/api"
	"APIGateway/internal/config"
)

type Server struct {
	API    *api.API
	Config *config.Config
}

func New() *Server {
	api := api.New()

	config := config.New()

	return &Server{
		API:    api,
		Config: config,
	}
}

func (s *Server) Run() {
	s.API.Run(s.Config.Port)
}
