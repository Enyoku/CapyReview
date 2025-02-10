package models

import "time"

type User struct {
	Email      string `json:"email" binding:"required,email"`
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	BIO        string
	Picture    string
	CreatedAt  time.Time
	LastOnline time.Time
}
