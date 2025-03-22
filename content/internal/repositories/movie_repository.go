package repositories

import (
	"contentService/internal/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MovieRepository определяет интерфейс для работы с фильмами
type MovieRepository interface {
	Create(ctx context.Context, movie *models.Movie) error
	GetByID(ctx context.Context, id string) (*models.Movie, error)
	Update(ctx context.Context, id string, movie *models.Movie) error
	Delete(ctx context.Context, id string) error
}

// Реализация интерфейса MovieRepository
type movieRepository struct {
	collection *mongo.Collection
}

// Создаёт новый экземпляр MovieRepository
func NewMovieRepository(client *mongo.Client) MovieRepository {
	return &movieRepository{
		collection: client.Database("entertainment").Collection("movies"),
	}
}

// Добавляет новый фильм в бд
func (r *movieRepository) Create(ctx context.Context, movie *models.Movie) error {
	_, err := r.collection.InsertOne(ctx, movie)
	return err
}

// Поиск по идентификатору фильма
func (r *movieRepository) GetByID(ctx context.Context, id string) (*models.Movie, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var movie models.Movie
	err = r.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&movie)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

// Обновляет фильм по идентификатору
func (r *movieRepository) Update(ctx context.Context, id string, movie *models.Movie) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": movie})
	return err
}

// Удаляет фильм по идентификатору
func (r *movieRepository) Delete(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectId})
	return err
}
