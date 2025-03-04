package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountGroup struct {
	group *gin.RouterGroup
}

func newAccountGroup(g *gin.RouterGroup) *AccountGroup {
	return &AccountGroup{group: g}
}

func (g *AccountGroup) RegisterRoutes() {
	g.group.GET("/", g.hiFunc)
}

func (api *AccountGroup) hiFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "hi"})
}
