package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID             string    `json:"id"`
	Firstname      string    `json:"firstname"`
	Lastname       string    `json:"lastname"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	HashedPasskey  string    `json:"hashed_passkey"`
	CreatedAt      time.Time `json:"created_at"`
}

type ActiveUser struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type SignedUser struct {
	ID       string
	Username string
	Email    string
	jwt.RegisteredClaims
}
