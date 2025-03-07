package server

import (
	"APIGateway/internal/api"
	"APIGateway/internal/config"
	"path/filepath"
)

// Path to find yaml config file
var yamlPath string = filepath.Join("internal", "config", "static-routes.yaml")

type Server struct {
	API    *api.API
	Config *config.Config
}

func New() *Server {

	config := config.New(yamlPath)

	api, err := api.New(config)
	if err != nil {

	}

	return &Server{
		API:    api,
		Config: config,
	}
}

func (s *Server) Run() {
	s.API.Run(s.Config.Env.Port)
}
