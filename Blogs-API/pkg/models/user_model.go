package models

import "time"

type User struct {
	ID             string `json:"id,omitempty"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	CreatedAt      time.Time
}
