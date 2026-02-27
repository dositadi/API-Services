package repository

import (
	m "blog/pkg/models"
	h "blog/pkg/utils"
)

func (db *BlogStore) LoginUser(user m.Login) (*m.ActiveUser, *m.ErrorMessage) {
	hashed_password, activeUser, err := h.GetHashedPassword(db.DB, user.Email)
	if err != nil {
		return nil, err
	}

	err2 := h.ComparePasswordAndHashed(hashed_password, user.Password)
	if err2 != nil {
		return nil, err2
	}
	return activeUser, nil
}
