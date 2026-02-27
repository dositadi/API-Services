package handlers

import (
	m "blog/pkg/models"
	h "blog/pkg/utils"
	"net/http"
)

func (b *BlogHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	var user m.Login

	if email == "" {
		errorMessage := h.ErrorMessageJson(h.EMPTY_FIELD, h.BAD_REQ_ERROR_CODE, h.EMAIL_EMPTY)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	}
	if password == "" {
		errorMessage := h.ErrorMessageJson(h.EMPTY_FIELD, h.BAD_REQ_ERROR_CODE, h.PASSWORD_EMPTY)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	}

	user.Email = email
	user.Password = password

	_, err := b.Store.LoginUser(user)
	if err != nil {
		errorMessage := h.ErrorMessageJson(h.INVALID_PASSWORD_ERROR, h.BAD_REQ_ERROR_CODE, h.INVALID_PASSWORD_ERROR_DETAIL)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
	}
}
