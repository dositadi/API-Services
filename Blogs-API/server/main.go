package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	home := Home{}
	//handler := BlogHandler{} // Use the constructor instead!

	router.HandleFunc("/", home.HomeHandler).Methods("GET")
	/* router.HandleFunc("/Blogs", handler.ListHandler).Methods("GET")
	router.HandleFunc("/Blogs/{id}", handler.GetHandler).Methods("GET")
	router.HandleFunc("/Blogs/{id}", handler.PatchHandler).Methods("PATCH")
	router.HandleFunc("/Blogs/{id}", handler.UpdateHandler).Methods("PUT")
	router.HandleFunc("/Blogs", handler.PostHandler).Methods("POST")
	router.HandleFunc("/Blogs/Delete", handler.DeleteHandler).Methods("DELETE") */

	server := &http.Server{
		Addr:              ":3000",
		Handler:           router,
		ReadTimeout:       500 * time.Millisecond,
		WriteTimeout:      500 * time.Millisecond,
		ReadHeaderTimeout: 500 * time.Millisecond,
		IdleTimeout:       1000 * time.Millisecond,
	}

	log.Fatal(server.ListenAndServe())
}
