package models

import (
	"errors"
	"time"
)

type Movie struct {
	Id          string    `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	ReleaseDate time.Time `json:"release_date" bson:"release_date"`
	Rating      float64   `json:"rating" bson:"rating"`
}

type MovieUpdate struct {
	Id          *string    `json:"id,omitempty"`
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	Rating      *float64   `json:"rating,omitempty"`
}

func (m *Movie) IsValid() (bool, error) {
	switch {
	case m.Title == "": // Проверяем, что название не пустое
		return false, errors.New("title cannot be empty")
	case m.Description == "": // Проверяем, что описание не пустое
		return false, errors.New("description cannot be empty")
	case m.ReleaseDate.After(time.Now()): // Проверяем, что дата выхода не в будущем
		return false, errors.New("release date cannot be in future")
	case m.Rating < 0 || m.Rating > 10: // Проверяем, что рейтинг находится в допустимом диапазоне
		return false, errors.New("rating must be between 0 and 10")
	}

	// Все проверки пройдены
	return true, nil
}
