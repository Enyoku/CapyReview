package api

import (
	"APIGateway/internal/api/middleware"
	"APIGateway/internal/config"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type API struct {
	router *gin.Engine
	config *config.Config
}

func New(config *config.Config) (*API, error) {
	// Initialize Gin
	gin.SetMode(gin.DebugMode) // TODO(change mode to Release)

	api := &API{
		router: gin.New(),
		config: config,
	}
	api.endpoints()
	return api, nil
}

func (api *API) Run(addr string) {
	api.router.Run(addr)
}

func (api *API) endpoints() {
	// Middleware
	api.router.Use(middleware.LoggerMiddleware())
	api.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:9001"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
	}))

	// Router Groups
	apiGroup := api.router.Group("/api")
	{
		accountGroup := newAccountGroup(apiGroup, api.config)
		accountGroup.RegisterRoutes()

		// Public Handlers
		apiGroup.GET("/ping", api.hiFunc)
	}
}

func (api *API) hiFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "pong"})
}
