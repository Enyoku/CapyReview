package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStringEnv(t *testing.T) {
	// Временная переменная окружения
	os.Setenv("TEST_KEY", "test_value")
	defer os.Unsetenv("TEST_KEY")

	// Переменная существует
	assert.Equal(t, "test_value", getStringEnv("TEST_KEY", "default_value"))

	// Переменная не существует
	assert.Equal(t, "default_value", getStringEnv("NON_EXIST_KEY", "default_value"))
}

func TestNewServiceConfig(t *testing.T) {
	yamlContent := `
services:
  auth_service:
    url: "http://localhost:9001"
    routes:
      - path: "/api/login"
        target: "/login"
        methods: ["POST"]
`

	// Создаём временный файл конфигурации
	tmpFile, err := os.CreateTemp("", "test-static-routes.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(yamlContent)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	// Читаем файл конфигурации
	serviceConfig, err := NewServiceConfig(tmpFile.Name())

	// Файл конфигурации прочитан успешно
	assert.NoError(t, err)
	assert.NotNil(t, serviceConfig)

	// Проверяем данные
	authService, exists := serviceConfig.Services["auth_service"]
	assert.True(t, exists)
	assert.Equal(t, "http://localhost:9001", authService.URL)

	// Проверяем маршруты
	assert.Len(t, authService.Routes, 1)

	// Проверяем первый маршрут
	assert.Equal(t, "/api/login", authService.Routes[0].Path)
	assert.Equal(t, "/login", authService.Routes[0].Target)
	assert.Equal(t, []string{"POST"}, authService.Routes[0].Methods)
}

func TestNewServiceConfig_FileNotFound(t *testing.T) {
	_, err := NewServiceConfig("non_existing_file.yaml")

	// Ошибка чтения файла
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read service config file")
}

func TestNew(t *testing.T) {
	yamlContent := `
services:
  auth_service:
    url: "http://localhost:9001"
    routes:
      - path: "/api/login"
        target: "/login"
        methods: ["POST"]
`

	tmpFile, err := os.CreateTemp("", "test-static-routes.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(yamlContent)
	if err != nil {
		t.Fatalf("Failed to write to temp file : %v", err)
	}
	tmpFile.Close()

	// Устанавливаем переменные окружения
	os.Setenv("port", ":9090")
	os.Setenv("secret_key", "my_secret_key")
	defer os.Unsetenv("port")
	defer os.Unsetenv("secret_key")

	// Создаём конфигурацию
	config := New(tmpFile.Name())

	// Конфигурация создана корректно
	assert.NotNil(t, config)
	assert.Equal(t, ":9090", config.Env.Port)
	assert.Equal(t, "my_secret_key", config.Env.SecretKey)
	assert.Contains(t, config.Services.Services, "auth_service")
	assert.Equal(t, "http://localhost:9001", config.Services.Services["auth_service"].URL)
}
