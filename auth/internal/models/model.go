package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	Id         int                `json:"id,omitempty"`
	Email      string             `json:"email" binding:"required,email"`
	Username   string             `json:"username" binding:"required"`
	Password   string             `json:"password" binding:"required"`
	BIO        string             `json:"bio,omitempty"`
	Picture    string             `json:"pic,omitempty"`
	CreatedAt  time.Time          `json:"created_at,omitempty"`
	LastOnline pgtype.Timestamptz `json:"last_online,omitempty"`
	Role       string             `json:"role"`
}

type UserProfileInfo struct {
	Email      string             `json:"email"`
	Username   string             `json:"username"`
	BIO        string             `json:"bio,omitempty"`
	Picture    string             `json:"pic,omitempty"`
	CreatedAt  time.Time          `json:"created_at,omitempty"`
	LastOnline pgtype.Timestamptz `json:"last_online,omitempty"`
}

type UserUpdate struct {
	Email      string             `json:"email,omitempty" binding:"omitempty,email"`
	Username   string             `json:"username,omitempty" binding:"omitempty"`
	Password   string             `json:"password,omitempty" binding:"omitempty"`
	BIO        string             `json:"bio,omitempty"`
	Picture    string             `json:"pic,omitempty"`
	LastOnline pgtype.Timestamptz `json:"last_online,omitempty"`
}

type LoginData struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// Decode JWT data
type Claims struct {
	UserId int
	Role   string
	jwt.RegisteredClaims
}
