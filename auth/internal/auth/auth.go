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
var ErrUnsupportedKey = errors.New("unsupported signing method")
var ErrTokenExpired = errors.New("token expired")
var ErrUnauthorized = errors.New("unauthorized")
var ErrTokenGeneration = errors.New("failed to generate token")

func GenerateToken(user *models.User, config config.JWT) (string, error) {

	claims := &models.Claims{
		UserId: user.Id,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.JWTExpiresIn)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.JWTSecretKey))
	if err != nil {
		return "", ErrTokenGeneration
	}
	return signedToken, nil
}

func GenerateRefreshToken(user *models.User, config config.JWT) (string, error) {
	claims := &models.Claims{
		UserId: user.Id,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.RefreshExpiresIn)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTSecretKey)
}

func ParseToken(tokenString string, config config.JWT) (*models.Claims, error) {

	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(config.JWTSecretKey), nil
	})
	if err != nil {
		log.Info().Msgf("%v", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		if claims.ExpiresAt.Unix() < time.Now().Unix() {
			return nil, ErrTokenExpired
		}
		return claims, nil
	}

	return nil, ErrInvalidToken
}
