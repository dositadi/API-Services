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

func (b *BlogStore) HealthCheck() *m.ErrorMessage {
	if err := b.DB.Ping(); err != nil {
		return &m.ErrorMessage{
			Error:   "Server Error",
			Details: []string{"Unable to connect to Server"},
			Code:    "500 Internal Server Error",
		}
	}
	return nil
}

func (db *BlogStore) List() ([]m.Blog, *m.ErrorMessage) {
	// TODO: Fuction Body
	return nil, nil
}

func (db *BlogStore) Get(id string) (m.Blog, *m.ErrorMessage) {
	// TODO: Fuction Body
	return m.Blog{}, nil
}

func (db *BlogStore) Post(blog m.Blog) *m.ErrorMessage {
	// TODO: Fuction Body
	return nil
}

func (db *BlogStore) Patch(id string, body m.Blog) ([]string, *m.ErrorMessage) {
	// TODO: Fuction Body
	return nil, nil
}

func (db *BlogStore) Update(id string, blog m.Blog) *m.ErrorMessage {
	// TODO: Fuction Body
	return nil
}

func (db *BlogStore) Delete(id string) *m.ErrorMessage {
	// TODO: Fuction Body
	return nil
}
