package server

import (
	"contentService/internal/repositories"
	"contentService/internal/services"

	"go.mongodb.org/mongo-driver/mongo"
)

// TODO(add a slice or struct of services and return it)
func initializeDependencies(db *mongo.Client) (services.MovieService, services.SeriesService, error) {
	//
	movieRepo := repositories.NewMovieRepository(db)
	seriesRepo := repositories.NewSeriesRepository(db)

	//
	movieService := services.NewMovieService(movieRepo)
	seriesService := services.NewSerialService(seriesRepo)

	return *movieService, *seriesService, nil
}
