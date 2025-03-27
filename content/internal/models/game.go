package models

import (
	"errors"
	"time"
)

type Game struct {
	Id          string    `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Publisher   string    `json:"publisher" bson:"publisher"`
	Developer   string    `json:"developer" bson:"developer"`
	ReleaseDate time.Time `json:"release_date" bson:"release_date"`
	Rating      float64   `json:"rating" bson:"rating"`
}

type GameUpdate struct {
	Id          *string    `json:"id,omitempty"`
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	Publisher   *string    `json:"publisher,omitempty"`
	Developer   *string    `json:"developer,omitempty"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	Rating      *float64   `json:"rating,omitempty"`
}

func (g *Game) IsValid() (bool, error) {
	switch {
	case g.Title == "":
		return false, errors.New("title cannot be empty")
	case g.Description == "":
		return false, errors.New("description cannot be empty")
	case g.Developer == "":
		return false, errors.New("developer cannot be empty")
	case g.ReleaseDate.After(time.Now()):
		return false, errors.New("release date cannot be in the future")
	case g.Rating < 0 || g.Rating > 10:
		return false, errors.New("rating must be between 0 and 10")
	}

	// Все проверки пройдены
	return true, nil
}
