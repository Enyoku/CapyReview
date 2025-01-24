package api

import (
	middleware "authService/internal/api/middlewares"
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
	// Initialize Gin
	gin.SetMode(gin.DebugMode)

	api := &API{
		router: gin.New(),
		db:     db,
	}
	api.Endpoints()
	return api, nil
}

func (api *API) Endpoints() {
	// Middlewares
	api.router.Use(middleware.HeaderMiddleware())
	api.router.Use(middleware.LoggerMiddleware())
	api.router.Use(gin.Recovery())

	// Handlers
	authGroup := api.router.Group("/api/account/")
	authGroup.GET("", api.HiFunc)
}

func (api *API) Run(addr string) {
	log.Printf("Starting server on port %v", addr)
	api.router.Run(addr)
}

func (api *API) HiFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
