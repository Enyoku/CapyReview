package api

import (
	"authService/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type API struct {
	router *gin.Engine
	db     *db.DB
}

func New(db *db.DB) (*API, error) {
	api := &API{
		router: gin.Default(),
		db:     db,
	}
	api.Endpoints()
	return api, nil
}

func (api *API) Endpoints() {
	api.router.Use(headerMiddleware())
	api.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func (api *API) Run(addr string) {
	log.Printf("Starting server on port %v", addr)
	api.router.Run(addr)
}
