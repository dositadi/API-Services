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
	if !h.ValidateEmail(email) {
		errorMessage := h.ErrorMessageJson(h.INVALID_EMAIL, h.BAD_REQ_ERROR_CODE, h.INVALID_EMAIL_DETAIL)
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

	activeUser, err := b.Store.LoginUser(user)
	if err != nil {
		if err.Error == h.SERVER_ERROR {
			errorMessage := h.ErrorMessageJson(h.SERVER_ERROR, h.SERVER_ERROR_CODE, h.SERVER_ERROR_DETAIL)
			h.Response(w, r, errorMessage, http.StatusInternalServerError)
			return
		} else if err.Error == h.PASS_MISMATCH_ERR {
			errorMessage := h.ErrorMessageJson(h.UNAUTHORIZED_ACCESS, h.UNAUTHORIZED_ACCESS_CODE, h.UNAUTHORIZED_ACCESS_DETAIL)
			h.Response(w, r, errorMessage, http.StatusUnauthorized)
			return
		} else {
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusBadRequest)
			return
		}
	}

	//  Generate access token
	accessToken, err2 := h.GenerateJWTAccessToken(*activeUser)
	if err2 != nil {
		errorMessage := h.ErrorMessageJson(h.UNAUTHORIZED_ACCESS, h.UNAUTHORIZED_ACCESS_CODE, h.UNAUTHORIZED_ACCESS_DETAIL)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	}

	// Generate refresh token
	refreshToken, err3 := h.GenerateJWTRefreshToken(*activeUser)
	if err3 != nil {
		errorMessage := h.ErrorMessageJson(h.UNAUTHORIZED_ACCESS, h.UNAUTHORIZED_ACCESS_CODE, h.UNAUTHORIZED_ACCESS_DETAIL)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{HttpOnly: true, Value: refreshToken, Name: "Refresh token"})

	h.Response(w, r, []byte(accessToken), http.StatusOK)
}
