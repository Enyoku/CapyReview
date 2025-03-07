package api

import (
	"APIGateway/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (g *AccountGroup) RegisterRoutes() {
	g.group.GET("/", g.hiFunc)
}

func (api *AccountGroup) hiFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "hi"})
}
