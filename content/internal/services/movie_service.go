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
	if _, err := movie.IsValid(); err != nil {
		return err
	} else {
		return s.repo.Create(ctx, movie)
	}
}

func (s *MovieService) GetMovie(ctx context.Context, id string) (*models.Movie, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *MovieService) UpdateMovie(ctx context.Context, id string, movie *models.MovieUpdate) error {
	// Получаем текущий фильм из бд
	existingMovie, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return errors.New("movie not found")
	}

	// Применяем обновления
	if movie.Title != nil {
		existingMovie.Title = *movie.Title
	}
	if movie.Description != nil {
		existingMovie.Description = *movie.Description
	}
	if movie.ReleaseDate != nil {
		existingMovie.ReleaseDate = *movie.ReleaseDate
	}
	if movie.Rating != nil {
		existingMovie.Rating = *movie.Rating
	}

	// Проверяем обновления
	if _, err := existingMovie.IsValid(); err != nil {
		return err
	}

	return s.repo.Update(ctx, id, existingMovie)
}

func (s *MovieService) DeleteMovie(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
