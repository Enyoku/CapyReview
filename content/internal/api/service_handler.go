package api

import (
	"contentService/internal/models"
	"contentService/internal/services"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SeriesHandler struct {
	service services.SeriesService
}

func NewSeriesHandler(service services.SeriesService) *SeriesHandler {
	return &SeriesHandler{service: service}
}

func (h *SeriesHandler) CreateSeries(c *gin.Context) {
	var series models.Series

	// Получаем тело запроса
	if err := c.ShouldBindJSON(&series); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Отправвляем полученный сериал в сервис
	if err := h.service.Create(context.Background(), &series); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create series"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "ok"})
}

func (h *SeriesHandler) GetSeriesById(c *gin.Context) {
	id := c.Param("id")

	val, err := h.service.GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "series not found"})
		return
	}

	c.JSON(http.StatusFound, val)
}

func (h *SeriesHandler) UpdateSeries(c *gin.Context) {
	var id string = c.Param("id")

	var updatedSeries models.SeriesUpdate

	if err := c.ShouldBindJSON(&updatedSeries); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	//
	if err := h.service.Update(context.Background(), id, &updatedSeries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update series"})
		return
	}

	c.JSON(http.StatusOK, updatedSeries)
}

func (h *SeriesHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.Delete(context.Background(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to delete series"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
