package handlers

import (
	m "blog/pkg/models"
	h "blog/pkg/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

const (
	EMAIL_EMPTY     = `Email should not be empty.`
	EMPTY_FIELD     = `Empty field.`
	FIRSTNAME_EMPTY = `Firstname should not be empty.`
	LASTNAME_EMPTY  = `Lastname should not be empty.`
	PASSWORD_EMPTY  = `Password should not be empty.`
	SHORTPASSWORD   = `Password should should be at least 6 characters.`
	PASSKEY_EMPTY   = `Passkey should not be empty.`
	LONGPASSKEY     = `Passkey should be at most 6 characters.`
	USERNAME_EMPTY  = `Username should not be empty.`

	SUCCESS_MESSAGE = `You have been registered successfully.`
)

func (b *BlogHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	firstname := mux.Vars(r)["firstname"]
	lastname := mux.Vars(r)["lastname"]
	password := mux.Vars(r)["password"]
	passkey := mux.Vars(r)["passkey"]
	username := mux.Vars(r)["username"]
	email := mux.Vars(r)["email"]
	var user m.User

	if email == "" {
		errorMessage := ErrorMessageJson(EMPTY_FIELD, "", EMAIL_EMPTY)
		w.Header().Set(CONTENT_TYPE, JSON)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	} else if firstname == "" {
		errorMessage := ErrorMessageJson(EMPTY_FIELD, "", FIRSTNAME_EMPTY)
		w.Header().Set(CONTENT_TYPE, JSON)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	} else if lastname == "" {
		errorMessage := ErrorMessageJson(EMPTY_FIELD, "", LASTNAME_EMPTY)
		w.Header().Set(CONTENT_TYPE, JSON)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	} else if username == "" {
		errorMessage := ErrorMessageJson(EMPTY_FIELD, "", USERNAME_EMPTY)
		w.Header().Set(CONTENT_TYPE, JSON)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	} else if password == "" {
		errorMessage := ErrorMessageJson(EMPTY_FIELD, "", PASSWORD_EMPTY)
		w.Header().Set(CONTENT_TYPE, JSON)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	} else if passkey == "" {
		errorMessage := ErrorMessageJson(EMPTY_FIELD, "", PASSKEY_EMPTY)
		w.Header().Set(CONTENT_TYPE, JSON)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	} else if len(passkey) > 6 {
		errorMessage := ErrorMessageJson(EMPTY_FIELD, "", LONGPASSKEY)
		w.Header().Set(CONTENT_TYPE, JSON)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	} else if len(password) < 6 {
		errorMessage := ErrorMessageJson(EMPTY_FIELD, "", SHORTPASSWORD)
		w.Header().Set(CONTENT_TYPE, JSON)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	}

	hashedPassword, err2 := h.SaltAndHashPassword(password)
	if err2 != nil {
		errorMessage := ErrorMessageJson(err2.Error, err2.Code, err2.Details...)
		w.Header().Set(CONTENT_TYPE, JSON)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	}

	hashedPasskey, err3 := h.SaltAndHashPassword(passkey)
	if err3 != nil {
		errorMessage := ErrorMessageJson(err3.Error, err3.Code, err3.Details...)
		w.Header().Set(CONTENT_TYPE, JSON)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
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
		errorMessage := ErrorMessageJson(err4.Error, err4.Code, err4.Details...)
		w.Header().Set(CONTENT_TYPE, JSON)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
		return
	}

	w.Header().Set(CONTENT_TYPE, JSON)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(SUCCESS_MESSAGE))
}
