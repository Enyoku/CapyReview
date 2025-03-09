package api

import (
	"APIGateway/internal/common"
	"APIGateway/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type AccountGroup struct {
	group  *gin.RouterGroup
	config *config.Config
}

func newAccountGroup(g *gin.RouterGroup, config *config.Config) *AccountGroup {
	return &AccountGroup{
		group:  g,
		config: config,
	}
}

func (a *AccountGroup) RegisterRoutes() {
	for _, route := range a.config.Services.Services["auth_service"].Routes {
		for _, method := range route.Methods {
			switch method {
			case "GET":
				a.group.GET(route.Path, func(c *gin.Context) {
					common.ProxyRequest(c, a.config, "auth_service", route.Target)
				})
			case "POST":
				a.group.POST(route.Path, func(c *gin.Context) {
					common.ProxyRequest(c, a.config, "auth_service", route.Target)
				})
			case "PATCH":
				a.group.PATCH(route.Path, func(c *gin.Context) {
					common.ProxyRequest(c, a.config, "auth_service", route.Target)
				})
			case "DELETE":
				a.group.DELETE(route.Path, func(c *gin.Context) {
					common.ProxyRequest(c, a.config, "auth_service", route.Target)
				})
			default:
				log.Error().Msgf("Unsupported method %s for route %s", method, route.Path)
			}
		}
	}
	a.group.GET("/", a.hiFunc)
}

func (a *AccountGroup) hiFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "hi"})
}
