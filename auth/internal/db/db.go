package db

import (
	"authService/internal/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
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

func AddUser(ctx context.Context, db *DB, u *models.User, hashedPassword string) (UserId int, err error) {
	var id int

	query := ` 
	 INSERT INTO users(email, username, password, bio, pic)
	 VALUES($1,$2,$3,$4,$5)
	 returning id
	`
	err = db.pool.QueryRow(ctx,
		query,
		u.Email, u.Username,
		hashedPassword, u.BIO,
		u.Picture).Scan(&id)

	if err != nil {
		log.Error().Msg(err.Error())
		return 0, nil
	}
	return id, nil
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

func FindUserByEmail(ctx context.Context, db *DB, email string) (models.User, bool) {
	var user models.User

	query := `
	SELECT id, email, username, pic, bio, created_at, last_online, role FROM users
	WHERE email = $1; 
	`

	err := db.pool.QueryRow(ctx, query, email).Scan(
		&user.Id,
		&user.Username,
		&user.BIO,
		&user.Picture,
		&user.CreatedAt,
		&user.LastOnline,
		&user.Role,
	)
	if err == pgx.ErrNoRows {
		return models.User{}, false
	} else if err != nil {
		log.Error().Msgf("Error querying database: %v", err)
		return models.User{}, false
	}
	return user, true
}

func FindUserByUsername(ctx context.Context, db *DB, username string) (models.User, bool) {
	var user models.User

	query := `
	SELECT id, email, username, pic, bio, created_at, last_online, role FROM users
	WHERE username = $1; 
	`

	err := db.pool.QueryRow(ctx, query, username).Scan(
		&user.Id,
		&user.Username,
		&user.BIO,
		&user.Picture,
		&user.CreatedAt,
		&user.LastOnline,
		&user.Role,
	)
	if err == pgx.ErrNoRows {
		return models.User{}, false
	} else if err != nil {
		log.Error().Msgf("Error querying database: %v", err)
		return models.User{}, false
	}
	return user, true
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
