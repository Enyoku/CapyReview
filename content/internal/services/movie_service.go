package services

import (
	"contentService/internal/models"
	"contentService/internal/repositories"
	"context"
	"errors"
)

type MovieService struct {
	repo repositories.MovieRepository
}

func NewMovieService(repo repositories.MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) CreateMovie(ctx context.Context, movie *models.Movie) error {
	return s.repo.Create(ctx, movie)
}

func (s *MovieService) GetMovie(ctx context.Context, id string) (*models.Movie, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *MovieService) UpdateMovie(ctx context.Context, id string, movie *models.Movie) error {
	//
	existingMovie, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return errors.New("movie not found")
	}

	//
	if movie.Title != "" {
		existingMovie.Title = movie.Title
	}
	if movie.Description != "" {
		existingMovie.Description = movie.Description
	}
	if !movie.ReleaseDate.IsZero() {
		existingMovie.ReleaseDate = movie.ReleaseDate
	}
	if movie.Rating != 0 {
		existingMovie.Rating = movie.Rating
	}

	//
	if err := s.repo.Update(ctx, id, existingMovie); err != nil {
		return errors.New("Failed to update movie")
	}

	return nil
}

func (s *MovieService) DeleteMovie(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
