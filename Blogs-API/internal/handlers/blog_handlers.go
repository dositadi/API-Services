package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	m "blog/pkg/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
)

type Home struct{}

func (h *Home) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is the home page!"))
}

func ErrorMessageJson(err string, code string, details ...string) []byte {
	errorMessage := m.ErrorMessage{
		Error:   err,
		Code:    code,
		Details: details,
	}

	errorJson, err2 := json.Marshal(errorMessage)
	if err2 != nil {
		fmt.Println(err2)
		return nil
	}
	return errorJson
}

func (b *BlogHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	errorMessage := ErrorMessageJson("Not Found", "404 Not Found", "The resource cannot be found.")
	w.Write(errorMessage)
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	errorMessage := ErrorMessageJson("Bad Request", "400 Bad Request", "The input is invalid.")
	w.Write(errorMessage)
}

func Conflict(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusConflict)
	errorMessage := ErrorMessageJson("Conflict", "409 Conflict", "The blog already exists.")
	w.Write(errorMessage)
}

/* type BlogPostInterface interface {
	List() ([]m.Blog, *m.ErrorMessage) /Blogs
	Post(blog m.Blog) *m.ErrorMessage /Blogs
	Patch(id string, field string, body any) *m.ErrorMessage /Blogs/{id}
	Update(id string, blog m.Blog) *m.ErrorMessage /Blogs/{id}
	Delete(id string) *m.ErrorMessage /Blogs/{id}
	HealthCheck() *m.ErrorMessage
} */

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
		b.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

func (b *BlogHandler) ListHandler(w http.ResponseWriter, r *http.Request) {
	blogs, err := b.Store.List()
	if err != nil {
		fmt.Println(err.Code)
		b.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

func (b *BlogHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	blog := m.Blog{}

	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rawID := uuid.NewString() //
	blog.Id = slug.Make(rawID)

	location := fmt.Sprintf("/Blogs/%s", blog.Id)
	/* tags := fmt.Sprintf("/Blogs/%s/Blog-Tag/%s", blog.Id, blog.Id)
	commentsLink := fmt.Sprintf("/Blogs/%s/Blog-Tag/%s", blog.Id,blog.Id) */

	blog.CommentCount = 0 // Todo: work on this!!!
	blog.PublishedAt = time.Now()
	//blog.Links = append(append(blog.Links, m.HyperLink{Relationship: "self", HyperReference: location}), m.HyperLink{Relationship: "Tags", HyperReference: tags})

	err2 := b.Store.Post(blog)
	if err2 != nil {
		if err2.Code == "400 Bad Request" {
			BadRequest(w, r)
			return
		} else {
			Conflict(w, r)
			return
		}
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
	blog := m.Blog{}

	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		fmt.Println(err)
	}

	field, err2 := b.Store.Patch(id, blog)
	if err2 != nil {
		b.NotFound(w, r)
		return
	}

	successMessage := fmt.Sprintf("%+v updated successfuly.", field)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(successMessage))
}
