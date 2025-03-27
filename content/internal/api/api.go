package api

import (
	"contentService/internal/middleware"
	"contentService/internal/services"

	"github.com/gin-gonic/gin"
)

type API struct {
	router        *gin.Engine
	movieService  services.MovieService
	seriesService services.SeriesService
	gameService   services.GameService
}

func New(movieService services.MovieService, seriesService services.SeriesService, gameService services.GameService) (*API, error) {
	gin.SetMode(gin.DebugMode) // TODO(change to release)
	api := &API{
		router:        gin.New(),
		movieService:  movieService,
		seriesService: seriesService,
		gameService:   gameService,
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
		v1.PATCH("/movies/:id", movieHandler.UpdateMovie)
		v1.DELETE("/movies/:id", movieHandler.DeleteMovie)

		// Series Handlers
		seriesHandler := NewSeriesHandler(a.seriesService)
		v1.POST("/series", seriesHandler.CreateSeries)
		v1.GET("/series/:id", seriesHandler.GetSeriesById)
		v1.PATCH("/series/:id", seriesHandler.UpdateSeries)
		v1.DELETE("/series/:id", seriesHandler.Delete)

		// Games Handlers
		gameHandler := NewGameHandler(a.gameService)
		v1.POST("/game", gameHandler.Create)
		v1.GET("/game/:id", gameHandler.GetByID)
		v1.PATCH("/game/:id", gameHandler.Update)
		v1.DELETE("/game/:id", gameHandler.Delete)
	}

}

func (a *API) Run(addr string) {
	a.router.Run(addr)
}
