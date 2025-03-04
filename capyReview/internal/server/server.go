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

	config := config.New()

	api, err := api.New(config)
	if err != nil {

	}

	return &Server{
		API:    api,
		Config: config,
	}
}

func (s *Server) Run() {
	s.API.Run(s.Config.Port)
}
