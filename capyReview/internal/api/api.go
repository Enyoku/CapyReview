package api

import (
	"APIGateway/internal/api/middleware"
	"APIGateway/internal/config"
	"net/http"

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

	// Router Groups
	public := api.router.Group("/")

	accountGroup := newAccountGroup(api.router.Group("/user"))
	accountGroup.RegisterRoutes()

	// Public Handlers
	public.GET("/", api.hiFunc)
}

func (api *API) hiFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "hi"})
}
