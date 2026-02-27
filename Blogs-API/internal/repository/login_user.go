package repository

import (
	m "blog/pkg/models"
	h "blog/pkg/utils"
	"database/sql"
)

const (
	GET_HASHED_QUERY = `SELECT hashed_password FROM users WHERE email=?`

	COMPARE_HASH_ERR        = `Incorrect password.`
	COMPARE_HASH_ERR_DETAIL = `You entered an incorrect password.`
)

func GetHashedPassword(db *sql.DB, email string) (string, *m.ErrorMessage) {
	row := db.QueryRow(GET_HASHED_QUERY, email)

	var hashed_password string

	err := row.Scan(&hashed_password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", &m.ErrorMessage{
				Error:   NOT_FOUND_ERR,
				Details: []string{NOT_FOUND_ERR_DETAIL},
				Code:    NOT_FOUND_ERR_CODE,
			}
		}
		return "", &m.ErrorMessage{
			Error:   CONN_ERR,
			Details: []string{err.Error()},
			Code:    "",
		}
	}
	return hashed_password, nil
}

func (db *BlogStore) LoginUser(email, password string) *m.ErrorMessage {
	hashed_password, err := GetHashedPassword(db.DB, email)
	if err != nil {
		return err
	}

	err2 := h.ComparePasswordAndHashed(hashed_password, password)
	if err2 != nil {
		return err2
	}
	return nil
}
