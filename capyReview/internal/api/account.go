package api

import (
	"APIGateway/internal/config"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

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
		a.group.Any(route.Path, func(c *gin.Context) {
			a.proxyRequest(c, "auth_service", route.Target)
		})
	}
	a.group.GET("/", a.hiFunc)
	a.group.POST("/login", a.loginHandler)
}

func (a *AccountGroup) hiFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "hi"})
}

func (a *AccountGroup) loginHandler(c *gin.Context) {
	a.proxyRequest(c, "auth_service", "/login")
}

func (a *AccountGroup) proxyRequest(c *gin.Context, serviceName, target string) {
	service, ok := a.config.Services.Services[serviceName]
	if !ok || service.URL == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Service %s not found dor URL is empty", serviceName)})
		return
	}

	var path string
	for _, route := range service.Routes {
		if route.Target == target {
			path = route.Path
			break
		}
	}

	if path == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Route not found for target" + target})
	}

	targetUrl := service.URL + path
	log.Info().Msgf("%v", targetUrl)

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	req, err := http.NewRequest(c.Request.Method, targetUrl, bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Faied to create request"})
		return
	}

	for k, v := range c.Request.Header {
		req.Header[k] = v
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to reach the target service"})
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to response body"})
		return
	}

	c.Data(resp.StatusCode, strings.Split(c.Request.Header.Get("Content-Type"), ";")[0], respBody)
}
