package main

import "net/http"

type Home struct{}

func (h *Home) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is the home page!"))
}
