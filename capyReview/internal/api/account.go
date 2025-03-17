package api

import (
	"APIGateway/internal/common"
	"APIGateway/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type AccountGroup struct {
	group  *gin.RouterGroup
	config *config.Config
	proxy  common.Proxy
}

func NewAccountGroup(g *gin.RouterGroup, config *config.Config, proxy common.Proxy) *AccountGroup {
	return &AccountGroup{
		group:  g.Group("/auth"),
		config: config,
		proxy:  proxy,
	}
}

func (a *AccountGroup) RegisterRoutes() {
	for _, route := range a.config.Services.Services["auth_service"].Routes {
		for _, method := range route.Methods {
			switch method {
			case "GET":
				a.group.GET(route.Target, func(c *gin.Context) {
					a.proxy.ProxyRequest(c, a.config, "auth_service", route.Target)
				})
			case "POST":
				a.group.POST(route.Target, func(c *gin.Context) {
					a.proxy.ProxyRequest(c, a.config, "auth_service", route.Target)
				})
			case "PATCH":
				a.group.PATCH(route.Target, func(c *gin.Context) {
					a.proxy.ProxyRequest(c, a.config, "auth_service", route.Target)
				})
			case "DELETE":
				a.group.DELETE(route.Target, func(c *gin.Context) {
					a.proxy.ProxyRequest(c, a.config, "auth_service", route.Target)
				})
			default:
				log.Error().Msgf("Unsupported method %s for route %s", method, route.Path)
			}
		}
	}
}
