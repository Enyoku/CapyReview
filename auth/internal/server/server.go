package server

import (
	"authService/internal/api"
	"authService/internal/config"
	"authService/internal/db"

	"github.com/rs/zerolog/log"
)

type Server struct {
	API    *api.API
	DB     *db.DB
	Config *config.Config
}

func New() (*Server, error) {

	config, err := config.New()
	if err != nil {
		log.Error().Msg(err.Error())
	}

	connString := "postgres://" + config.DB.Username + ":" + config.DB.Password + "@" + config.DB.Host + "/" + config.DB.Name

	db, err := db.New(connString)
	if err != nil {
		log.Error().Msg(err.Error())
	}

	api, err := api.New(db)
	if err != nil {
		log.Error().Msg(err.Error())
	}

	s := &Server{
		API:    api,
		DB:     db,
		Config: config,
	}
	return s, nil
}

func (s *Server) Run() {
	s.API.Run(s.Config.Port)
}
