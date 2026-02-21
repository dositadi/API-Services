package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	m "blog/models"

	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
)

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

func NotFound(w http.ResponseWriter, r *http.Request) {
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
} */

func (b *BlogHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	rawID := mux.Vars(r)["id"]
	id := slug.Make(rawID)

	blog, err := b.Store.Get(id)
	if err != nil {
		NotFound(w, r)
		return
	}

	blogJson, err2 := json.Marshal(blog)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(blogJson)
}

func (b *BlogHandler) ListHandler(w http.ResponseWriter, r *http.Request) {
	blogs, err := b.Store.List()
	if err != nil {
		fmt.Println(err.Code)
		NotFound(w, r)
		return
	}

	blogsJson, err2 := json.Marshal(blogs)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(blogsJson)
}

func (b *BlogHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	blog := m.Blog{}

	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Fallback to this.
	//rawID := uuid //
	//blog.Id = slug.Make(rawID)

	location := fmt.Sprintf("/Blogs/%s", blog.Id)

	blog.CommentCount = len(blog.Comments)
	blog.PublishedAt = time.RFC1123Z
	blog.Links = append(blog.Links, m.HyperLink{Relationship: "self", HyperReference: location})

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
		NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	//w.Header()
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
		NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
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
		NotFound(w, r)
		return
	}

	successMessage := fmt.Sprintf("%+v updated successfuly.", field)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(successMessage))
}
