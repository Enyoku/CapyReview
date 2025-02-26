package server

import (
	"APIGateway/internal/api"
	"APIGateway/internal/config"
	"APIGateway/internal/db"
)

type Server struct {
	API    *api.API
	DB     *db.DB
	Config *config.Config
}

func New() *Server {
	db := db.New()

	api := api.New(db)

	config := config.New()

	return &Server{
		API:    api,
		DB:     db,
		Config: config,
	}
}

func (s *Server) Run() {
	// s.API.Run(s.Config)
}
