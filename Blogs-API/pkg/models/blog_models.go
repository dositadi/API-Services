package models

type Blog struct {
	Id           string      `json:"id,omitempty"`
	Title        string      `json:"title"`
	Content      string      `json:"content"`
	Author       string      `json:"author"`
	PublishedAt  string      `json:"published_at,omitempty"`
	Archive      bool        `json:"archive,omitempty"`
	Tags         []string    `json:"tags,omitempty"`
	Comments     []Comment   `json:"comments,omitempty"`
	CommentCount int         `json:"comment_count,omitempty"`
	Links        []HyperLink `json:"links,omitempty"`
}

type HyperLink struct {
	Relationship   string `json:"rel"`
	HyperReference string `json:"href"`
}

type Comment struct {
	Id        string      `json:"id,omitempty"`
	PostID    string      `json:"post_id"`
	Author    string      `json:"author"`
	Content   string      `json:"content"`
	CreatedAt string      `json:"created_at,omitempty"`
	Links     []HyperLink `json:"links,omitempty"`
}
