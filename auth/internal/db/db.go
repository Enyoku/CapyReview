package db

import (
	"authService/internal/models"
	"context"
	"fmt"

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

func GetUserInfo(ctx context.Context, db *DB, id int) (*models.UserProfileInfo, error) {
	var user models.UserProfileInfo

	query := `
	SELECT email, username, bio, pic, created_at, last_online FROM users
	WHERE id = $1;
	`

	err := db.pool.QueryRow(ctx, query, id).Scan(&user.Email,
		&user.Username, &user.BIO, &user.Picture, &user.CreatedAt, &user.LastOnline)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUser(ctx context.Context, db *DB, id int) error {
	query := `
	DELETE FROM users
	WHERE id = $1;
	`

	res, err := db.pool.Exec(ctx, query, id)
	if err != nil {
		log.Error().Msgf("%v", err)
		return err
	}

	rowsAffected := res.RowsAffected()

	if rowsAffected == 0 {
		log.Error().Msgf("user with ID: %v not found", id)
		return fmt.Errorf("user with ID: %v not found", id)
	}

	return nil
}

// func updateUserInfo()
