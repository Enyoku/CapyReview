package server

import (
	"contentService/internal/repositories"
	"contentService/internal/services"

	"go.mongodb.org/mongo-driver/mongo"
)

func initializeDependencies(db *mongo.Client) (services.MovieService, error) {
	//
	movieRepo := repositories.NewMovieRepository(db)

	//
	movieService := services.NewMovieService(movieRepo)

	return *movieService, nil
}
