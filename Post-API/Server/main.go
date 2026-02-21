package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"posts"
	m "posts"
	"strconv"

	"github.com/gorilla/mux"
)

type Home struct{}

func (h Home) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is the Home Page!"))
}

type postsMethods interface {
	Get(id string) (m.Post, error)
	Put(id string, post m.Post) error
	List() (map[string]m.Post, error)
	Delete(id string) error
	Update(id string, post m.Post) error
}

type PostServerHandler struct {
	Store postsMethods
}

func NewPostHandler(s postsMethods) *PostServerHandler {
	return &PostServerHandler{
		Store: s,
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found!"))
}

func NoContent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("No Content!"))
}

func StatusFailed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusFailedDependency)
	w.Write([]byte("Action Failed!"))
}

func MediaContentNotSupported(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnsupportedMediaType)
	w.Write([]byte("Unsupported Media Type!"))
}

func (h *PostServerHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	post, err := h.Store.Get(id)
	if err != nil {
		fmt.Println(err)
		NotFound(w, r)
		return
	}

	json, err2 := json.Marshal(post)
	if err2 != nil {
		NoContent(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func (h *PostServerHandler) PutHandler(w http.ResponseWriter, r *http.Request) {
	random := rand.Intn(500)
	id := ""
	if random < 10 {
		id = "00" + strconv.Itoa(random)
	} else if random < 100 {
		id = "0" + strconv.Itoa(random)
	} else {
		id = strconv.Itoa(random)
	}

	post := m.Post{}

	if r.Header.Get("Content-Type") != "application/json" {
		MediaContentNotSupported(w, r)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		NoContent(w, r)
		return
	}

	err2 := h.Store.Put(id, post)
	if err2 != nil {
		fmt.Println(err2)
		StatusFailed(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post Added"))
}

func (h *PostServerHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.Store.Delete(id)
	if err != nil {
		fmt.Println(err)
		NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post Deleted"))
}

func (h *PostServerHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	post := m.Post{}

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		NoContent(w, r)
		return
	}

	err2 := h.Store.Update(id, post)
	if err2 != nil {
		fmt.Println(err2)
		NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Action Completed!"))
}

func (h *PostServerHandler) ListHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Store.List()
	if err != nil {
		fmt.Println(err)
		NoContent(w, r)
		return
	}

	postsJson, err2 := json.Marshal(posts)
	if err2 != nil {
		NoContent(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(postsJson)
}

func main() {
	router := mux.NewRouter()
	home := Home{}

	postStore := posts.NewPostStore("/home/ositadinma/L2E Fellowship/Personal/Post-API/posts.json")
	postServerHandler := NewPostHandler(postStore)

	router.HandleFunc("/", home.HomeHandler).Methods("GET")
	router.HandleFunc("/Posts/{id}", postServerHandler.DeleteHandler).Methods("DELETE")
	router.HandleFunc("/Posts", postServerHandler.PutHandler).Methods("POST")
	router.HandleFunc("/Posts/{id}", postServerHandler.GetHandler).Methods("GET")
	router.HandleFunc("/Posts/{id}", postServerHandler.UpdateHandler).Methods("PUT")
	router.HandleFunc("/Posts", postServerHandler.ListHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
