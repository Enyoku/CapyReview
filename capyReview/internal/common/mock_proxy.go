package common

import (
	"APIGateway/internal/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockProxy struct {
	mock.Mock
}

func (m *MockProxy) ProxyRequest(c *gin.Context, cfg *config.Config, serviceName, target string) {
	log.Print("called")
	m.Called(c, cfg, serviceName, target)
}
