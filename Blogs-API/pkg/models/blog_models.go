package models

import "time"

type Blog struct {
	Id           string    `json:"id"`
	UserID       string    `json:"user_id"` // Foreign key
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	PublishedAt  time.Time `json:"published_at,omitempty"`
	Archive      bool      `json:"archive,omitempty"`
	CommentCount int       `json:"comment_count,omitempty"`
}

type Tag struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Junction table for blog tags.
type BlogTags struct {
	TagID  string `json:"tag_id"`
	BlogID string `json:"blog_id"` // Foreign key
}

type Links struct {
	Id             string `json:"id"`
	BlogID         string `json:"blog_id"` // Foreign key
	Relationship   string `json:"rel"`
	HyperReference string `json:"href"`
}

type Comment struct {
	Id        string    `json:"id"`
	BlogID    string    `json:"blog_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
