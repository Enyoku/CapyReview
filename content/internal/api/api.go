package api

import (
	"contentService/internal/middleware"
	"contentService/internal/services"

	"github.com/gin-gonic/gin"
)

type API struct {
	router       *gin.Engine
	movieService services.MovieService
}

func New(movieService services.MovieService) (*API, error) {
	gin.SetMode(gin.DebugMode) // TODO(change to release)
	api := &API{
		router:       gin.New(),
		movieService: movieService,
	}
	api.Endpoints()
	return api, nil
}

func (a *API) Endpoints() {
	// Middlewares
	a.router.Use(middleware.LoggingMiddleware())

	a.router.Use(gin.Recovery())

	v1 := a.router.Group("/api/v1")
	{
		// Movie Handlers
		movieHandler := NewMovieHandler(a.movieService)
		v1.POST("/movies", movieHandler.CreateMovie)
		v1.GET("/movies/:id", movieHandler.GetMovie)
	}

}

func (a *API) Run(addr string) {
	a.router.Run(addr)
}
