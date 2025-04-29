package server

import "reviewService/internal/api"

type Server struct {
	api *api.API
	// DB
	// config *config.Config
}

func NewServer() (*Server, error) {
	api, err := api.NewAPI()
	if err != nil {

	}

	return &Server{
		api: api,
	}, nil
}

func (s *Server) Run(addr string) {
	s.api.Run(addr)
}
