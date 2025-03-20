package api

import (
	"contentService/internal/middleware"

	"github.com/gin-gonic/gin"
)

type API struct {
	router *gin.Engine
}

func New() (*API, error) {
	gin.SetMode(gin.DebugMode) // TODO(change to release)
	api := &API{
		router: gin.New(),
	}
	api.Endpoints()
	return api, nil
}

func (a *API) Endpoints() {
	// Middlewares
	a.router.Use(middleware.LoggingMiddleware())

	a.router.Use(gin.Recovery())

	// Handlers

}

func (a *API) Run(addr string) {
	a.router.Run(addr)
}
