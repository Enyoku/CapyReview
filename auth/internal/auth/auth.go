package auth

import (
	"authService/internal/config"
	"authService/internal/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

var ErrInvalidToken = errors.New("invalid_token")
var ErrUnauthorized = errors.New("unauthorized")

func GenerateToken(user *models.User, config config.JWT) (string, error) {
	claims := &models.Claims{
		UserId: user.Id, // TODO(в бд сериал, поэтому id в модели изначально пустой)
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.JWTExpiresIn)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(config.JWTSecretKey))
}

func GenerateRefreshToken(user *models.User, config config.JWT) (string, error) {
	claims := &models.Claims{
		UserId: user.Id, // TODO(в бд сериал, поэтому id в модели изначально пустой)
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.RefreshExpiresIn)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(config.JWTSecretKey)
}

func ParseToken(tokenString string, config config.JWT) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(config.JWTSecretKey), nil
	})
	if err != nil {
		log.Error().Msgf("%v", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*models.Claims); !ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}
