package utils

import "net/http"

func Response(w http.ResponseWriter, r *http.Request, message []byte, code int) {
	w.WriteHeader(code)
	w.Header().Set(CONTENT_TYPE, JSON)
	w.Write(message)
}

