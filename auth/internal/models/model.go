package models

import "time"

type User struct {
	Email      string    `json:"email" binding:"required,email"`
	Username   string    `json:"username" binding:"required"`
	Password   string    `json:"password" binding:"required"`
	BIO        string    `json:"bio,omitempty"`
	Picture    string    `json:"pic,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	LastOnline time.Time `json:"last_online,omitempty"`
}

type UserProfileInfo struct {
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	BIO        string    `json:"bio,omitempty"`
	Picture    string    `json:"pic,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	LastOnline time.Time `json:"last_online,omitempty"`
}

type Uid struct {
	Id int `json:"id"`
}

// Id         int       `json:"-,omitempty"`
