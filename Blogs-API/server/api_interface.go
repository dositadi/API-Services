package main

import m "blog/models"

type BlogPostInterface interface {
	List() ([]m.Blog, *m.ErrorMessage)
	Get(id string) (m.Blog, *m.ErrorMessage)
	Post(blog m.Blog) *m.ErrorMessage
	Patch(id string, body m.Blog) ([]string, *m.ErrorMessage)
	Update(id string, blog m.Blog) *m.ErrorMessage
	Delete(id string) *m.ErrorMessage
}

type BlogHandler struct {
	Store BlogPostInterface
}

func NewBlogHandler(store BlogPostInterface) *BlogHandler {
	return &BlogHandler{
		Store: store,
	}
}
