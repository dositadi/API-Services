package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	h "blog/pkg/utils"
)

func (b *BlogHandler) ListCommentsHandler(w http.ResponseWriter, r *http.Request) {
	blog_id := mux.Vars(r)["blog_id"]

	ctx := r.Context()

	comments, err := b.Store.ListComments(ctx, blog_id)
	if err != nil {
		if err.Error == h.SERVER_ERROR {
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusInternalServerError)
			return
		} else {
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set(h.CONTENT_TYPE, h.JSON)
	w.WriteHeader(http.StatusOK)
	err2 := json.NewEncoder(w).Encode(comments)
	if err2 != nil {
		errorMessage := h.ErrorMessageJson(h.SERVER_ERROR, h.SERVER_ERROR_CODE, h.SERVER_ERROR_DETAIL)
		h.Response(w, r, errorMessage, http.StatusInternalServerError)
		return
	}
}

func (b *BlogHandler) GetCommentHandler(w http.ResponseWriter, r *http.Request) {
	blog_id := mux.Vars(r)["blog_id"]
	id := mux.Vars(r)["id"]

	ctx := r.Context()

	comment, err := b.Store.GetComment(ctx, blog_id, id)
	if err != nil {
		switch err.Error {
		case h.SERVER_ERROR:
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusInternalServerError)
			return
		case h.NOT_FOUND_ERR:
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusNotFound)
			return
		default:
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set(h.CONTENT_TYPE, h.JSON)
	w.WriteHeader(http.StatusOK)
	if err2 := json.NewEncoder(w).Encode(comment); err2 != nil {
		errorMessage := h.ErrorMessageJson(h.SERVER_ERROR, h.SERVER_ERROR_CODE, h.SERVER_ERROR_DETAIL)
		h.Response(w, r, errorMessage, http.StatusInternalServerError)
		return
	}
}
