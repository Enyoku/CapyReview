package common

import (
	"APIGateway/internal/config"
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Proxy interface {
	ProxyRequest(c *gin.Context, cfg *config.Config, serviceName, target string)
}

// Реализация интерфейса
type DefaultProxy struct{}

func (p *DefaultProxy) ProxyRequest(c *gin.Context, cfg *config.Config, serviceName, target string) {
	// Получаем описание сервиса из конфигурации
	service, ok := cfg.Services.Services[serviceName]
	if !ok || service.URL == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Service %s not foundd or URL is empty" + serviceName})
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

	// Формируем полный URL
	targetUrl := service.URL + path

	// Читаем тело запроса
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// Восстанавливаем тело для дальнейшего использования
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	// Новый HTTP-запрос
	req, err := http.NewRequest(c.Request.Method, targetUrl, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// Копируем заголовки из входящего запроса
	for k, v := range c.Request.Header {
		req.Header[k] = v
	}

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to reach the target service"})
		return
	}
	defer resp.Body.Close()

	// Пересылаем Cookie из ответа сервиса
	for _, cookie := range resp.Cookies() {
		c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	}

	// Читаем ответ от сервиса
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	// Возвращаем ответ клиенту
	c.Data(resp.StatusCode, strings.Split(c.Request.Header.Get("Content-Type"), ";")[0], respBody)
}
