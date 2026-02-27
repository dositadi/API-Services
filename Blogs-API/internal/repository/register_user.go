package repository

import (
	m "blog/pkg/models"
	"database/sql"
)

const (
	INSERT_USER_QUERY = `INSERT INTO users (id, first_name,last_name,username,email,hashed_password,hashed_passkey) VALUES (?,?,?,?,?,?,?)`
)

func (db *BlogStore) RegisterUser(user m.User) *m.ErrorMessage {
	_, err := db.DB.Exec(INSERT_USER_QUERY, user.ID, user.Firstname, user.Lastname, user.Username, user.Email, user.HashedPassword, user.HashedPasskey)
	if err != nil {
		if err == sql.ErrConnDone {
			return &m.ErrorMessage{
				Error:   CONN_ERR,
				Details: []string{err.Error()},
				Code:    "",
			}
		}
		return &m.ErrorMessage{
			Error:   INSERTION_ERR,
			Details: []string{err.Error()},
		}
	}
	return nil
}
