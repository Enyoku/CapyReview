package services

import (
	"contentService/internal/models"
	"contentService/internal/repositories"
	"context"
)

type SeriesService struct {
	repo repositories.SeriesRepository
}

func NewSerialService(repo repositories.SeriesRepository) *SeriesService {
	return &SeriesService{repo: repo}
}

func (s *SeriesService) Create(ctx context.Context, series *models.Series) error {
	// Проверяем полученные данные
	if _, err := series.IsValid(); err != nil {
		return s.repo.Create(ctx, series)
	} else {
		return err
	}
}

func (s *SeriesService) GetByID(ctx context.Context, id string) (*models.Series, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *SeriesService) Update(ctx context.Context, id string, series *models.SeriesUpdate) error {
	// Получаем текущий сериал из бд
	existingSeries, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return err
	}

	// Проверяем, корректные ли данные
	if series.Title != nil {
		existingSeries.Title = *series.Title
	}
	if series.Description != nil {
		existingSeries.Description = *series.Description
	}
	if series.ReleaseDate != nil {
		existingSeries.ReleaseDate = *series.ReleaseDate
	}
	if series.Seasons != nil {
		existingSeries.Seasons = *series.Seasons
	}
	if series.Rating != nil {
		existingSeries.Rating = *series.Rating
	}

	// Проверяем изменения
	if _, err := existingSeries.IsValid(); err != nil {
		return err
	}

	return s.repo.Update(context.Background(), id, existingSeries)
}

func (s *SeriesService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(context.Background(), id)
}
