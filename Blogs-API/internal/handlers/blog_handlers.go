package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "blog/pkg/models"
	h "blog/pkg/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
)

const (
	CONTENT_TYPE = "Content-Type"
	JSON         = "application/json"
)

type Home struct{}

func (h *Home) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is the home page!"))
}

func (b *BlogHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	errorMessage := h.ErrorMessageJson("Not Found", "404", "The resource cannot be found.")
	w.Write(errorMessage)
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	errorMessage := h.ErrorMessageJson("Bad Request", "400", "The input is invalid.")
	w.Write(errorMessage)
}

func Conflict(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusConflict)
	errorMessage := h.ErrorMessageJson("Conflict", "409", "The blog already exists.")
	w.Write(errorMessage)
}

func (b *BlogHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if err := b.Store.HealthCheck(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errMessage := fmt.Sprintf("%s:%+v\n", err.Error, err.Details)
		w.Write([]byte(errMessage))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("System is alright!\n"))
}

func (b *BlogHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	rawID := mux.Vars(r)["id"]
	id := slug.Make(rawID)

	blog, err := b.Store.Get(id)
	if err != nil {
		if err.Error == h.NOT_FOUND_ERR {
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusNotFound)
			return
		}
		errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
		h.Response(w, r, errorMessage, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

func (b *BlogHandler) ListHandler(w http.ResponseWriter, r *http.Request) {
	blogs, err := b.Store.List()
	if err != nil {
		if err.Error == h.CONN_ERR {
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusNotFound)
			return
		}

		if err.Error == h.SERVER_ERROR {
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusNotFound)
			return
		}
		errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
		h.Response(w, r, errorMessage, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

func (b *BlogHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	content := mux.Vars(r)["content"]

	var blog m.Blog

	/* err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		err := m.ErrorMessage{
			Error:   h.BAD_REQ_ERROR,
			Details: []string{h.BAD_REQ_ERROR_DETAILS},
			Code:    h.BAD_REQ_ERROR_CODE,
		}
		errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	}
	*/

	blog.Title = &title
	blog.Content = &content
	blog.Id = uuid.NewString()
	blog.Archive = false
	blog.CommentCount = 0

	location := fmt.Sprintf("/blogs/%s", blog.Id)

	err2 := b.Store.Post(blog)
	if err2 != nil {
		errorMessage := h.ErrorMessageJson(err2.Error, err2.Code, err2.Details...)
		h.Response(w, r, errorMessage, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", location)
	w.Write([]byte("Blog created"))
}

func (b *BlogHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	rawID := mux.Vars(r)["id"]
	id := slug.Make(rawID)

	err := b.Store.Delete(id)
	if err != nil {
		b.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Blog deleted successfully."))
}

func (b *BlogHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	rawID := mux.Vars(r)["id"]
	id := slug.Make(rawID)
	blog := m.Blog{}

	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		fmt.Println(err)
	}

	err2 := b.Store.Update(id, blog)
	if err2 != nil {
		b.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Blog updated successfully."))
}

func (b *BlogHandler) PatchHandler(w http.ResponseWriter, r *http.Request) {
	rawID := mux.Vars(r)["id"]
	id := slug.Make(rawID)
	var query map[string]string

	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		fmt.Println(err)
	}

	err2 := b.Store.Patch(id, query)
	if err2 != nil {
		b.NotFound(w, r)
		return
	}

	successMessage := "Updated successfuly."

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(successMessage))
}
