package services

import (
	"contentService/internal/models"
	"contentService/internal/repositories"
	"context"
	"errors"
)

type SeriesService struct {
	repo repositories.SeriesRepository
}

func NewSerialService(repo repositories.SeriesRepository) (*SeriesService, error) {
	return &SeriesService{repo: repo}, nil
}

func (s *SeriesService) Create(series models.Series) error {
	// Проверяем полученные данные
	if series.IsValid() {
		return s.repo.Create(context.Background(), series)
	} else {
		return errors.New("invalid credentials")
	}
}

func (s *SeriesService) GetByID(id string) (*models.Series, error) {
	return s.repo.GetByID(context.Background(), id)
}

func (s *SeriesService) Update(id string, series models.SeriesUpdate) error {
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
	if !existingSeries.IsValid() {
		return errors.New("failed to update series")
	}

	return s.repo.Update(context.Background(), id, *existingSeries)
}

func (s *SeriesService) Delete(id string) error {
	return s.repo.Delete(context.Background(), id)
}
