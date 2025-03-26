package models

import (
	"errors"
	"time"
)

type Season struct {
	Number      int       `json:"number,omitempty" bson:"number"`
	Title       string    `json:"title" bson:"title"`
	ReleaseDate time.Time `json:"release_date" bson:"release_date"`
	Rating      float64   `json:"rating" bson:"rating"`
	Episodes    int       `json:"episodes" bson:"episodes"`
}

type Series struct {
	Id          string    `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Seasons     []Season  `json:"seasons,omitempty" bson:"seasons,omitempty"`
	ReleaseDate time.Time `json:"release_date" bson:"release_date"`
	Rating      float64   `json:"rating" bson:"rating"`
}

type SeriesUpdate struct {
	Id          *string    `json:"id,omitempty"`
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	Seasons     *[]Season  `json:"seasons,omitempty"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	Rating      *float64   `json:"rating,omitempty"`
}

type SeasonUpdate struct {
	Number      *int       `json:"number,omitempty"`
	Title       *string    `json:"title,omitempty"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	Rating      *float64   `json:"rating,omitempty"`
	Episodes    *int       `json:"episodes,omitempty"`
}

func (s *Series) IsValid() (bool, error) {
	switch {
	case s.Title == "": // Проверяем, что название не пустое
		return false, errors.New("series titile cannot be empty")
	case s.Description == "": // Проверяем, что описание не пустое
		return false, errors.New("series description cannot be empty")
	case s.ReleaseDate.After(time.Now()): // Проверяем, что дата выхода не в будущем
		return false, errors.New("series release date cannot be in the future")
	case s.Rating < 0 || s.Rating > 10: // Проверяем, что рейтинг находится в допустимом диапазоне
		return false, errors.New("series rating must be between 0 and 10")
	}

	// Проверяем каждый сезон
	for _, season := range s.Seasons {
		season.IsValid()
	}

	// Все проверки пройдены
	return true, nil
}

func (s *Season) IsValid() (bool, error) {
	switch {
	case s.Title == "": // Проверяем, что название не пустое
		return false, errors.New("season titile cannot be empty")
	case s.ReleaseDate.After(time.Now()): // Проверяем, что дата выхода не в будущем
		return false, errors.New("season release date cannot be in the future")
	case s.Rating < 0 || s.Rating > 10: // Проверяем, что рейтинг находится в допустимом диапазоне
		return false, errors.New("season rating must be between 0 and 10")
	}

	// Все проверки пройдены
	return true, nil
}
