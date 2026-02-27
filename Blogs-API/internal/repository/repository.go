package repository

import (
	"database/sql"
	"strings"

	m "blog/pkg/models"
	h "blog/pkg/utils"

	"github.com/google/uuid"
)

type BlogStore struct {
	DB *sql.DB
}

func NewBlogStore(db *sql.DB) *BlogStore {
	return &BlogStore{
		DB: db,
	}
}

var (
	User = &m.ActiveUser{}
)

func CurrentUser(user m.ActiveUser) {
	*User = user
}

func (b *BlogStore) HealthCheck() *m.ErrorMessage {
	if err := b.DB.Ping(); err != nil {
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{h.SERVER_ERROR_DETAIL},
			Code:    h.SERVER_ERROR,
		}
	}
	return nil
}

func (db *BlogStore) List() ([]m.Blog, *m.ErrorMessage) {
	rows, err := db.DB.Query(h.LIST_QUERY)
	if err != nil {
		return nil, &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{err.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	defer rows.Close()

	var blogs []m.Blog

	for rows.Next() {
		var blog m.Blog
		if err := rows.Scan(&blog.Id, &blog.UserID, &blog.Title, &blog.Content, &blog.PublishedAt, &blog.Archive, &blog.CommentCount); err != nil {
			if err == sql.ErrConnDone {
				return blogs, &m.ErrorMessage{
					Error:   h.CONN_ERR,
					Details: []string{err.Error()},
					Code:    h.SERVER_ERROR_CODE,
				}
			}
			return blogs, &m.ErrorMessage{
				Error:   h.SERVER_ERROR,
				Details: []string{err.Error()},
				Code:    h.SERVER_ERROR_CODE,
			}
		}
		blogs = append(blogs, blog)
	}

	if err := rows.Err(); err != nil {
		return blogs, &m.ErrorMessage{
			Error:   h.ROW_SCAN_ERR,
			Details: []string{err.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}
	return blogs, nil
}

func (db *BlogStore) Get(id string) (m.Blog, *m.ErrorMessage) {
	row := db.DB.QueryRow(h.GET_QUERY, id)

	var blog m.Blog
	err := row.Scan(&blog.Id, &blog.UserID, &blog.Title, &blog.Content, &blog.PublishedAt, &blog.Archive, &blog.CommentCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return m.Blog{}, &m.ErrorMessage{
				Error:   h.NOT_FOUND_ERR,
				Details: []string{h.NOT_FOUND_ERR_DETAIL},
				Code:    h.NOT_FOUND_ERR_CODE,
			}
		}
		return m.Blog{}, &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{err.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}
	return blog, nil
}

func (db *BlogStore) Post(blog m.Blog) *m.ErrorMessage {
	blog.UserID = ""

	_, err := db.DB.Exec(h.POST_QUERY, blog.Id, blog.UserID, blog.Title, blog.Content, blog.Archive, blog.CommentCount)
	if err != nil {
		return &m.ErrorMessage{
			Error:   h.INSERTION_ERR,
			Details: []string{err.Error()},
			Code:    "",
		}
	}

	selfLink := m.Links{
		Id:             uuid.NewString(),
		BlogID:         blog.Id,
		Relationship:   "Self",
		HyperReference: "/blogs/" + blog.Id,
	}

	commentsLink := m.Links{
		Id:             uuid.NewString(),
		BlogID:         blog.Id,
		Relationship:   "Comments",
		HyperReference: "/blogs/" + blog.Id + "/comment",
	}

	_, err2 := db.DB.Exec(h.LINK_INSERT_QUERY, selfLink.Id, selfLink.BlogID, selfLink.Relationship, selfLink.HyperReference, commentsLink.Id, commentsLink.BlogID, commentsLink.Relationship, commentsLink.HyperReference)
	if err2 != nil {
		return &m.ErrorMessage{
			Error:   h.INSERTION_ERR,
			Details: []string{err2.Error()},
			Code:    "",
		}
	}
	return nil
}

func (db *BlogStore) Patch(id string, query map[string]string) *m.ErrorMessage { // Amend its parameters
	blog, err := db.Get(id)
	if err != nil {
		return err
	}

	for k := range query {
		switch k {
		case "title":
			title := query[k]
			blog.Title = &title
		case "content":
			content := query[k]
			blog.Content = &content
		case "archive":
			switch strings.ToLower(query[k]) {
			case "true":
				blog.Archive = true
			case "false":
				blog.Archive = false
			}
		}
	}

	err2 := db.Update(id, blog)
	if err2 != nil && err2.Error == h.CONN_ERR {
		return err2
	}
	return nil
}

func (db *BlogStore) Update(id string, blog m.Blog) *m.ErrorMessage {
	result, err := db.DB.Exec(h.UPDATE_QUERY, blog.Title, blog.Content, blog.PublishedAt, blog.CommentCount, id)
	if err != nil {
		return &m.ErrorMessage{
			Error:   h.CONN_ERR,
			Details: []string{err.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	rows_affected, err2 := result.RowsAffected()
	if err2 != nil {
		return &m.ErrorMessage{
			Error:   h.CONN_ERR,
			Details: []string{err.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	if rows_affected == 0 {
		return &m.ErrorMessage{
			Error:   h.NOT_FOUND_ERR,
			Details: []string{h.NOT_FOUND_ERR_DETAIL},
			Code:    h.NOT_FOUND_ERR_CODE,
		}
	}
	return nil
}

func (db *BlogStore) Delete(id string) *m.ErrorMessage {
	result, err := db.DB.Exec(h.DELETE_QUERY, id)
	if err != nil {
		return &m.ErrorMessage{
			Error:   h.CONN_ERR,
			Details: []string{err.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	rowsAffected, err2 := result.RowsAffected()
	if err2 != nil {
		return &m.ErrorMessage{
			Error:   h.CONN_ERR,
			Details: []string{err2.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	if rowsAffected == 0 {
		return &m.ErrorMessage{
			Error:   h.NOT_FOUND_ERR,
			Details: []string{h.NOT_FOUND_ERR_DETAIL},
			Code:    h.NOT_FOUND_ERR_CODE,
		}
	}
	return nil
}
