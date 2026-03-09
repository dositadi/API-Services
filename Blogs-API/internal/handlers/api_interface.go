package handlers

import (
	m "blog/pkg/models"
	"context"
)

type BlogPostInterface interface {
	// Blog functions
	List() ([]m.Blog, *m.ErrorMessage)
	Get(id string) (m.Blog, *m.ErrorMessage)
	Post(blog m.Blog) *m.ErrorMessage
	Patch(id string, query map[string]string) *m.ErrorMessage
	Update(id string, blog m.Blog) *m.ErrorMessage
	Delete(id string) *m.ErrorMessage

	// Comments functions
	ListComments(ctx context.Context, blog_id string) ([]m.Comment, *m.ErrorMessage)
	GetComment(ctx context.Context, blog_id, id string) (m.Comment, *m.ErrorMessage)
	PostComment(ctx context.Context, blog_id, id string, comment m.Comment) *m.ErrorMessage
	PatchComment(ctx context.Context, blog_id, id string, query map[string]string) *m.ErrorMessage
	UpdateComment(ctx context.Context, blog_id, id string, comment m.Comment) *m.ErrorMessage
	DeleteComment(ctx context.Context, blog_id, id string) *m.ErrorMessage

	// Health Check
	HealthCheck() *m.ErrorMessage

	// Auth Functions
	RegisterUser(user m.User) *m.ErrorMessage
	LoginUser(user m.Login) (*m.ActiveUser, *m.ErrorMessage)
}

type BlogHandler struct {
	Store BlogPostInterface
}

func NewBlogHandler(store BlogPostInterface) *BlogHandler {
	return &BlogHandler{
		Store: store,
	}
}
