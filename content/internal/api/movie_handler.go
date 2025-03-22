package api

import (
	"contentService/internal/models"
	"contentService/internal/services"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	service services.MovieService
}

func NewMovieHandler(service services.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var movie models.Movie

	//
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	//
	if err := h.service.CreateMovie(context.Background(), &movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create movie"})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

func (h *MovieHandler) GetMovie(c *gin.Context) {
	id := c.Param("id")

	//
	movie, err := h.service.GetMovie(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie

	//
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	//
	if err := h.service.UpdateMovie(context.Background(), id, &movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
		return
	}

	//
	c.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	//
	if err := h.service.DeleteMovie(context.Background(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete movie"})
		return
	}

	//
	c.JSON(http.StatusNoContent, nil)
}
