package models

import "time"

type User struct {
	ID             string `json:"id"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	HashedPasskey  string `json:"hashed_passkey"`
	CreatedAt      time.Time
}

type ActiveUser struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}
