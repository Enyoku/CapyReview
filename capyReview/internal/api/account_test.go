package api

import (
	"APIGateway/internal/common"
	"APIGateway/internal/config"
	"APIGateway/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterRoutes(t *testing.T) {
	cfg := &config.Config{
		Services: &config.ServiceConfig{
			Services: map[string]models.Service{
				"auth_service": {
					URL: "http://localhost:9001",
					Routes: []models.Route{
						{
							Path:    "/api/login",
							Target:  "/login",
							Methods: []string{"POST"},
						},
					},
				},
			},
		},
	}

	// Создание mock-объекта
	mockProxy := new(common.MockProxy)
	mockProxy.On("ProxyRequest", mock.Anything, mock.MatchedBy(func(config *config.Config) bool {
		return config.Services.Services["auth_service"].URL == "http://localhost:9001"
	}), "auth_service", "/login").Return()

	// Создание тестого роутера
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Создание группы роутеров аутентификации с mock-прокси
	accountGroup := NewAccountGroup(router.Group("/"), cfg, mockProxy)

	// Регистрация маршрутов
	accountGroup.RegisterRoutes()

	req, _ := http.NewRequest("POST", "/auth/login", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	// Проверка, что mock был вызван
	mockProxy.AssertNumberOfCalls(t, "ProxyRequest", 1)
	mockProxy.AssertCalled(t, "ProxyRequest", mock.Anything, cfg, "auth_service", "/login")
	assert.Equal(t, http.StatusOK, resp.Code)
}
