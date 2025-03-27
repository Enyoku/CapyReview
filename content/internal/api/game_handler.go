package api

import (
	"contentService/internal/models"
	"contentService/internal/services"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	service *services.GameService
}

func NewGameHandler(service services.GameService) *GameHandler {
	return &GameHandler{service: &service}
}

func (h *GameHandler) Create(c *gin.Context) {
	var game models.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.service.Create(context.Background(), &game); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create a new game"})
		return
	}

	c.JSON(http.StatusCreated, game)
}

func (h *GameHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	game, err := h.service.GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "game not found"})
		return
	}

	c.JSON(http.StatusOK, game)
}

func (h *GameHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var updatedGame models.GameUpdate

	if err := c.ShouldBindJSON(updatedGame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.service.Update(context.Background(), id, &updatedGame); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update game"})
		return
	}

	c.JSON(http.StatusOK, updatedGame)
}

func (h *GameHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.Delete(context.Background(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to delete game"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
