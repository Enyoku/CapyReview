package models

import "time"

type User struct {
	Email      string
	Username   string
	Password   string
	BIO        string
	Picture    string
	CreatedAt  time.Time
	LastOnline time.Time
}
