package models

import "time"

type Season struct {
	Number      int       `json:"number,omitempty" bson:"number"`
	Title       string    `json:"title" bson:"title"`
	ReleaseDate time.Time `json:"release_date" bson:"release_date"`
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
