package repositories

import (
	"contentService/internal/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SeriesRepository interface {
	Create(ctx context.Context, series models.Series) error
	GetByID(ctx context.Context, id string) (*models.Series, error)
	Update(ctx context.Context, id string, series models.Series) error
	Delete(ctx context.Context, id string) error
}

type seriesRepository struct {
	collection *mongo.Collection
}

func NewSeriesRepository(client *mongo.Client) SeriesRepository {
	return &seriesRepository{
		collection: client.Database("entertainment").Collection("series"),
	}
}

func (s *seriesRepository) Create(ctx context.Context, series models.Series) error {
	_, err := s.collection.InsertOne(ctx, series)
	return err
}

func (s *seriesRepository) GetByID(ctx context.Context, id string) (*models.Series, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var series models.Series
	err = s.collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&series)
	if err != nil {
		return nil, err
	}

	return &series, nil
}

func (s *seriesRepository) Update(ctx context.Context, id string, series models.Series) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.collection.UpdateByID(ctx, bson.M{"_id": objId}, series)
	return err
}

func (s *seriesRepository) Delete(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.collection.DeleteOne(ctx, bson.M{"_id": objId})
	return err
}
