package repository

import (
	m "blog/pkg/models"
	h "blog/pkg/utils"
	"database/sql"
)

func (db *BlogStore) RegisterUser(user m.User) *m.ErrorMessage {
	var exists bool
	row := db.DB.QueryRow(h.CHECK_USER_QUERY, user.Email, user.Username)

	err := row.Scan(&exists)
	if err != nil {
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{h.SERVER_ERROR_DETAIL},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	if exists {
		return &m.ErrorMessage{
			Error:   h.ALREADY_EXISTS_ERROR,
			Details: []string{h.ALREADY_EXISTS_ERROR_DETAIL},
			Code:    h.ALREADY_EXISTS_ERROR_CODE,
		}
	}

	_, err2 := db.DB.Exec(h.INSERT_USER_QUERY, user.ID, user.Firstname, user.Lastname, user.Username, user.Email, user.HashedPassword, user.HashedPasskey)
	if err2 != nil {
		if err2 == sql.ErrConnDone {
			return &m.ErrorMessage{
				Error:   h.SERVER_ERROR,
				Details: []string{h.SERVER_ERROR_DETAIL},
				Code:    h.SERVER_ERROR_CODE,
			}
		}
		return &m.ErrorMessage{
			Error:   h.INSERTION_ERR,
			Details: []string{err2.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}
	return nil
}
