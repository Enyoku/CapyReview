package models

import "time"

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

func (m *Movie) IsValid() bool {
	// Проверяем, что название не пустое
	if m.Title == "" {
		return false
	}

	// Проверяем, что описание не пустое
	if m.Description == "" {
		return false
	}

	// Проверяем, что дата выхода не в будущем
	if m.ReleaseDate.After(time.Now()) {
		return false
	}

	// Проверяем, что рейтинг находится в допустимом диапазоне
	if m.Rating < 0 || m.Rating > 10 {
		return false
	}

	// Все проверки пройдены
	return true
}
