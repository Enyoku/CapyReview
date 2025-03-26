package repositories

import (
	"contentService/internal/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GameRepository определяет интерфейс для работы с бд
type GameRepository interface {
	Create(ctx context.Context, game *models.Game) error
	GetByID(ctx context.Context, id string) (*models.Game, error)
	Update(ctx context.Context, id string, updatedGame *models.Game) error
	Delete(ctx context.Context, id string) error
}

// Реализация интерфейса GameRepository
type gameRepository struct {
	collection *mongo.Collection
}

// Создаёт новый экземпляр GameRepository
func NewGameRepository(client *mongo.Client) GameRepository {
	return &gameRepository{
		collection: client.Database("entertainment").Collection("games"),
	}
}

// Создаёт запись новой игры в бд
func (r *gameRepository) Create(ctx context.Context, game *models.Game) error {
	_, err := r.collection.InsertOne(ctx, game)
	return err
}

// Поиск по идентификатору игры
func (r *gameRepository) GetByID(ctx context.Context, id string) (*models.Game, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var game models.Game
	err = r.collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

// Обновляет фильм по идентификатору
func (r *gameRepository) Update(ctx context.Context, id string, updatedGame *models.Game) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": updatedGame})
	return err
}

// Удаляет фильм по идентификатору
func (r *gameRepository) Delete(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objId})
	return err
}
