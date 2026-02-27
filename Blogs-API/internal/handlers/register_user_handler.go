package handlers

import (
	m "blog/pkg/models"
	h "blog/pkg/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (b *BlogHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user m.User

	email := r.PostFormValue("email")
	firstname := r.PostFormValue("firstname")
	lastname := r.PostFormValue("lastname")
	username := r.PostFormValue("userpassword")
	passkey := r.PostFormValue("passkey")
	password := r.PostFormValue("password")

	if email == "" {
		errorMessage := h.ErrorMessageJson(h.EMPTY_FIELD, h.BAD_REQ_ERROR_CODE, h.EMAIL_EMPTY)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	} else if firstname == "" {
		errorMessage := h.ErrorMessageJson(h.EMPTY_FIELD, h.BAD_REQ_ERROR_CODE, h.FIRSTNAME_EMPTY)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	} else if lastname == "" {
		errorMessage := h.ErrorMessageJson(h.EMPTY_FIELD, h.BAD_REQ_ERROR_CODE, h.LASTNAME_EMPTY)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	} else if username == "" {
		errorMessage := h.ErrorMessageJson(h.EMPTY_FIELD, h.BAD_REQ_ERROR_CODE, h.USERNAME_EMPTY)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	} else if password == "" {
		errorMessage := h.ErrorMessageJson(h.EMPTY_FIELD, h.BAD_REQ_ERROR_CODE, h.PASSWORD_EMPTY)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	} else if passkey == "" {
		errorMessage := h.ErrorMessageJson(h.EMPTY_FIELD, h.BAD_REQ_ERROR_CODE, h.PASSKEY_EMPTY)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	} else if len(passkey) > 6 {
		errorMessage := h.ErrorMessageJson(h.EMPTY_FIELD, h.BAD_REQ_ERROR_CODE, h.LONGPASSKEY)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	} else if len(password) < 6 {
		errorMessage := h.ErrorMessageJson(h.EMPTY_FIELD, h.BAD_REQ_ERROR_CODE, h.SHORTPASSWORD)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	}

	hashedPassword, err2 := h.SaltAndHashPassword(password)
	if err2 != nil {
		errorMessage := h.ErrorMessageJson(err2.Error, err2.Code, err2.Details...)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	}

	hashedPasskey, err3 := h.SaltAndHashPassword(passkey)
	if err3 != nil {
		errorMessage := h.ErrorMessageJson(err3.Error, err3.Code, err3.Details...)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	}

	// Validate email with regex.

	user.ID = uuid.NewString()
	user.CreatedAt = time.Now()
	user.Email = email
	user.Firstname = firstname
	user.Lastname = lastname
	user.Username = username
	user.HashedPassword = hashedPassword
	user.HashedPasskey = hashedPasskey

	err4 := b.Store.RegisterUser(user)
	if err4 != nil {
		errorMessage := h.ErrorMessageJson(err4.Error, err4.Code, err4.Details...)
		if err4.Code == h.SERVER_ERROR_CODE {
			h.Response(w, r, errorMessage, http.StatusInternalServerError)
			return
		}
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	}

	h.Response(w, r, []byte(h.SUCCESS_MESSAGE), http.StatusOK)
}
