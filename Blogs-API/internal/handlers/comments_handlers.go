package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	m "blog/pkg/models"
	h "blog/pkg/utils"
)

func (b *BlogHandler) ListCommentsHandler(w http.ResponseWriter, r *http.Request) {
	blog_id := mux.Vars(r)["blog_id"]

	ctx := r.Context()

	comments, err := b.Store.ListComments(ctx, blog_id)
	if err != nil {
		errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
		h.Response(w, r, errorMessage, http.StatusInternalServerError)
		return
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

func (b *BlogHandler) PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	content := r.PostFormValue("content")
	blog_id := mux.Vars(r)["blog_id"]

	ctx := r.Context()

	var comment m.Comment

	id := uuid.NewString()

	comment.BlogID = blog_id
	comment.Id = id
	comment.Content = content

	err := b.Store.PostComment(ctx, blog_id, id, comment)
	if err != nil {
		switch err.Error {
		case h.SERVER_ERROR:
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusInternalServerError)
			return
		case h.CONN_ERR:
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusInternalServerError)
			return
		case h.NOT_FOUND_ERR:
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusNotFound)
			return
		}
	}

	w.Header().Set(h.CONTENT_TYPE, h.JSON)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(h.SUCCESS_MESSAGE))
}

func (b *BlogHandler) UpdateCommentHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["comment_id"]
	content := r.PostFormValue("content")
	blog_id := mux.Vars(r)["blog_id"]

	ctx := r.Context()

	var comment m.Comment

	comment.BlogID = blog_id
	comment.Id = id
	comment.Content = content

	err := b.Store.UpdateComment(ctx, blog_id, id, comment)
	if err != nil {
		switch err.Error {
		case h.SERVER_ERROR:
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusInternalServerError)
			return
		case h.CONN_ERR:
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusInternalServerError)
			return
		case h.NOT_FOUND_ERR:
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusNotFound)
			return
		}
	}

	w.Header().Set(h.CONTENT_TYPE, h.JSON)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(h.SUCCESS_MESSAGE))
}

func (b *BlogHandler) PatchCommentHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["comment_id"]
	content := r.PostFormValue("content")
	blog_id := mux.Vars(r)["blog_id"]

	ctx := r.Context()

	query := map[string]string{
		"content": content,
	}

	err := b.Store.PatchComment(ctx, blog_id, id, query)
	if err != nil {
		switch err.Error {
		case h.SERVER_ERROR:
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusInternalServerError)
			return
		case h.CONN_ERR:
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusInternalServerError)
			return
		case h.NOT_FOUND_ERR:
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusNotFound)
			return
		}
	}

	w.Header().Set(h.CONTENT_TYPE, h.JSON)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(h.SUCCESS_MESSAGE))
}

func (b *BlogHandler) DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	blog_id := mux.Vars(r)["blog_id"]
	id := mux.Vars(r)["id"]

	ctx := r.Context()

	err := b.Store.DeleteComment(ctx, blog_id, id)
	if err != nil {
		switch err.Error {
		case h.SERVER_ERROR:
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusInternalServerError)
			return
		case h.CONN_ERR:
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusInternalServerError)
			return
		case h.NOT_FOUND_ERR:
			errorMessage := h.ErrorMessageJson(err.Error, err.Code, err.Details...)
			h.Response(w, r, errorMessage, http.StatusNotFound)
			return
		}
	}

	w.Header().Set(h.CONTENT_TYPE, h.JSON)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(h.SUCCESS_MESSAGE))
}
