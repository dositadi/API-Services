package utils

import (
	m "blog/pkg/models"
	"database/sql"
)

func GetHashedPassword(db *sql.DB, email string) (string, *m.ActiveUser, *m.ErrorMessage) {
	row := db.QueryRow(GET_HASHED_QUERY, email)

	var id, firstname, lastname, username, hashed_password string

	err := row.Scan(&id, &firstname, &lastname, &username, &hashed_password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil, &m.ErrorMessage{
				Error:   PASS_MISMATCH_ERR,
				Details: []string{PASS_MISMATCH_ERR_DETAIL},
				Code:    "",
			}
		}
		return "", nil, &m.ErrorMessage{
			Error:   SERVER_ERROR,
			Details: []string{SERVER_ERROR_DETAIL},
			Code:    SERVER_ERROR_CODE,
		}
	}

	user := &m.ActiveUser{
		ID:        id,
		Username:  username,
		Email:     email,
	}
	return hashed_password, user, nil
}
