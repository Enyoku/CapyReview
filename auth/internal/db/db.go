package db

import (
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
