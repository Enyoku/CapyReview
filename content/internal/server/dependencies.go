package server

import (
	"contentService/internal/repositories"
	"contentService/internal/services"

	"go.mongodb.org/mongo-driver/mongo"
)

// TODO(add a slice or struct of services and return it)
func initializeDependencies(db *mongo.Client) (services.MovieService, services.SeriesService, services.GameService, error) {
	//
	movieRepo := repositories.NewMovieRepository(db)
	seriesRepo := repositories.NewSeriesRepository(db)
	gameRepo := repositories.NewGameRepository(db)

	//
	movieService := services.NewMovieService(movieRepo)
	seriesService := services.NewSerialService(seriesRepo)
	gameService := services.NewGameService(gameRepo)

	return *movieService, *seriesService, *gameService, nil
}
