package db

import (
	"authService/internal/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type DB struct {
	pool *pgxpool.Pool
}

func New(connString string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	return &DB{
		pool: pool,
	}, err
}

func AddUser(ctx context.Context, db *DB, u *models.User, hashedPassword string) error {
	var id int

	query := ` 
	 INSERT INTO users(email, username, password, bio, pic)
	 VALUES($1,$2,$3,$4,$5)
	 returning id
	`
	err := db.pool.QueryRow(ctx,
		query,
		u.Email, u.Username,
		hashedPassword, u.BIO,
		u.Picture).Scan(&id)

	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}
	return nil
}
