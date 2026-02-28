package utils

import (
	m "blog/pkg/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTAccessToken(user m.ActiveUser) (string, error) {
	var secretKey = os.Getenv("JWT_SECRET_KEY")

	claims := &m.SignedUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(1 * time.Hour))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return result, nil
}

func GenerateJWTRefreshToken(user m.ActiveUser) (string, error) {
	var secretKey = os.Getenv("JWT_SECRET_KEY")

	claims := &m.SignedUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(168 * time.Hour))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return result, nil
}
