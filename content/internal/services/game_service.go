package services

import (
	"contentService/internal/models"
	"contentService/internal/repositories"
	"context"
)

type GameService struct {
	repo repositories.GameRepository
}

func NewGameService(repo repositories.GameRepository) *GameService {
	return &GameService{repo: repo}
}

func (g *GameService) Create(ctx context.Context, game *models.Game) error {
	if _, err := game.IsValid(); err != nil {
		return err
	} else {
		return g.repo.Create(ctx, game)
	}
}

func (g *GameService) GetByID(ctx context.Context, id string) (*models.Game, error) {
	return g.repo.GetByID(ctx, id)
}

func (g *GameService) Update(ctx context.Context, id string, udpatedGame *models.GameUpdate) error {
	existedGame, err := g.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// Применяем обновления
	if udpatedGame.Title != nil {
		existedGame.Title = *udpatedGame.Title
	}
	if udpatedGame.Description != nil {
		existedGame.Description = *udpatedGame.Description
	}
	if udpatedGame.Publisher != nil {
		existedGame.Publisher = *udpatedGame.Publisher
	}
	if udpatedGame.Developer != nil {
		existedGame.Developer = *udpatedGame.Developer
	}
	if udpatedGame.ReleaseDate != nil {
		existedGame.ReleaseDate = *udpatedGame.ReleaseDate
	}
	if udpatedGame.Rating != nil {
		existedGame.Rating = *udpatedGame.Rating
	}

	if _, err := existedGame.IsValid(); err != nil {
		return err
	} else {
		return g.repo.Update(ctx, id, existedGame)
	}
}

func (g *GameService) Delete(ctx context.Context, id string) error {
	return g.repo.Delete(ctx, id)
}
