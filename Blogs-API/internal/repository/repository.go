package repository

import (
	"database/sql"

	m "blog/pkg/models"
)

type BlogStore struct {
	DB *sql.DB
}

func NewBlogStore(db *sql.DB) *BlogStore {
	return &BlogStore{
		DB: db,
	}
}

const (
	LIST_QUERY   = `SELECT * FROM blogs ORDER BY published_at DESC`
	GET_QUERY    = `SELECT id, user_id, title, content, published_at, archive, comment_count FROM blogs WHERE id=?`
	POST_QUERY   = `INSERT INTO blogs (id, user_id, title, content,archive,comment_count) VALUES (?,?,?,?,?,?)`
	UPDATE_QUERY = `UPDATE blogs SET title=?,content=?,published_at=?, comment_count=? WHERE id=?`
	DELETE_QUERY = `DELETE FROM blogs WHERE id=?`
)

const (
	CONN_ERR             = `Connection Error.`
	NOT_FOUND_ERR        = `Not Found.`
	NOT_FOUND_ERR_CODE   = `404`
	NOT_FOUND_ERR_DETAIL = `Blog was not found on the database.`
	SERVER_ERROR         = `Internal server error.`
	SERVER_ERROR_CODE    = `500`
	SERVER_ERROR_DETAIL  = `Unable to connect to Server.`
	DELETE_ERROR         = `Deletion error.`
	ROW_SCAN_ERR         = `Row scan error`
	INSERTION_ERR        = `Insertion error.`
)

func (b *BlogStore) HealthCheck() *m.ErrorMessage {
	if err := b.DB.Ping(); err != nil {
		return &m.ErrorMessage{
			Error:   SERVER_ERROR,
			Details: []string{SERVER_ERROR_DETAIL},
			Code:    SERVER_ERROR,
		}
	}
	return nil
}

func (db *BlogStore) List() ([]m.Blog, *m.ErrorMessage) {
	rows, err := db.DB.Query(LIST_QUERY)
	if err != nil {
		return nil, &m.ErrorMessage{
			Error:   SERVER_ERROR,
			Details: []string{err.Error()},
			Code:    "",
		}
	}

	defer rows.Close()

	var blogs []m.Blog

	for rows.Next() {
		var blog m.Blog
		if err := rows.Scan(&blog.Id, &blog.UserID, &blog.Title, &blog.Content, &blog.PublishedAt, &blog.Archive, &blog.CommentCount); err != nil {
			if err == sql.ErrConnDone {
				return blogs, &m.ErrorMessage{
					Error:   CONN_ERR,
					Details: []string{err.Error()},
					Code:    "",
				}
			}
			return blogs, &m.ErrorMessage{
				Error:   "",
				Details: []string{err.Error()},
				Code:    "",
			}
		}
		blogs = append(blogs, blog)
	}

	if err := rows.Err(); err != nil {
		return blogs, &m.ErrorMessage{
			Error:   ROW_SCAN_ERR,
			Details: []string{err.Error()},
			Code:    "",
		}
	}
	return blogs, nil
}

func (db *BlogStore) Get(id string) (m.Blog, *m.ErrorMessage) {
	row := db.DB.QueryRow(GET_QUERY, id)

	var blog m.Blog
	err := row.Scan(&blog.Id, &blog.UserID, &blog.Title, &blog.Content, &blog.PublishedAt, &blog.Archive, &blog.CommentCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return m.Blog{}, &m.ErrorMessage{
				Error:   NOT_FOUND_ERR,
				Details: []string{NOT_FOUND_ERR_DETAIL},
				Code:    NOT_FOUND_ERR_CODE,
			}
		}
		return m.Blog{}, &m.ErrorMessage{
			Error:   SERVER_ERROR,
			Details: []string{err.Error()},
			Code:    SERVER_ERROR_CODE,
		}
	}
	return blog, nil
}

func (db *BlogStore) Post(blog m.Blog) *m.ErrorMessage {
	_, err := db.DB.Exec(POST_QUERY, blog.Id, blog.UserID, blog.Title, blog.Content, blog.Archive, blog.CommentCount)
	if err != nil {
		return &m.ErrorMessage{
			Error:   INSERTION_ERR,
			Details: []string{err.Error()},
			Code:    "",
		}
	}
	return nil
}

func (db *BlogStore) Patch(id string, query map[string]string) (int, *m.ErrorMessage) { // Amend its parameters
	blog, err := db.Get(id)
	if err != nil {
		return 0, err
	}

	for k := range query {
		switch k {
		case "title":
			blog.Title = query[k]
		case "content":
			blog.Content = query[k]
		case "archive":
			switch query[k] {
			case "true":
				blog.Archive = true
			case "false":
				blog.Archive = false
			}
		}
	}

	rows_affected, err2 := db.Update(id, blog)
	if err2 != nil {
		return 0, err2
	}
	return rows_affected, nil
}

func (db *BlogStore) Update(id string, blog m.Blog) (int, *m.ErrorMessage) {
	if _, err := db.DB.Exec(UPDATE_QUERY, blog.Title, blog.Content, blog.PublishedAt, blog.CommentCount, id); err != nil {
		if err == sql.ErrNoRows {
			return 0, &m.ErrorMessage{
				Error:   NOT_FOUND_ERR,
				Details: []string{NOT_FOUND_ERR_DETAIL},
				Code:    NOT_FOUND_ERR_CODE,
			}
		}
		return 0, &m.ErrorMessage{
			Error:   CONN_ERR,
			Details: []string{err.Error()},
			Code:    "",
		}
	}
	return 0, nil
}

func (db *BlogStore) Delete(id string) *m.ErrorMessage {
	if _, err := db.DB.Exec(DELETE_QUERY, id); err != nil {
		return &m.ErrorMessage{
			Error:   DELETE_ERROR,
			Details: []string{err.Error()},
			Code:    "",
		}
	}
	return nil
}
