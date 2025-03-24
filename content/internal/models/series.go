package models

import "time"

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

func (s *Series) IsValid() bool {
	// Проверяем, что название не пустое
	if s.Title == "" {
		return false
	}

	// Проверяем, что описание не пустое
	if s.Description == "" {
		return false
	}

	// Проверяем, что дата выхода не в будущем
	if s.ReleaseDate.After(time.Now()) {
		return false
	}

	// Проверяем, что рейтинг находится в допустимом диапазоне
	if s.Rating < 0 || s.Rating > 10 {
		return false
	}

	// Проверяем каждый сезон
	for _, season := range s.Seasons {
		season.IsValid()
	}

	// Все проверки пройдены
	return true
}

func (s *Season) IsValid() bool {
	// Проверяем, что название не пустое
	if s.Title == "" {
		return false
	}

	// Проверяем, что дата выхода не в будущем
	if s.ReleaseDate.After(time.Now()) {
		return false
	}

	// Проверяем, что рейтинг находится в допустимом диапазоне
	if s.Rating < 0 || s.Rating > 10 {
		return false
	}

	// Все проверки пройдены
	return true
}
