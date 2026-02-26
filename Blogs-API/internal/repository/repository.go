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
	rows, err := db.DB.Query(`SELECT * FROM blogs`)
	if err != nil {
		return nil, &m.ErrorMessage{
			Error:   "Database access error.",
			Details: []string{"Could not fetch data from the database."},
			Code:    "",
		}
	}

	defer rows.Close()

	var blogs []m.Blog

	for rows.Next() {
		var blog m.Blog
		if err := rows.Scan(&blog.Id, &blog.UserID, &blog.Title, &blog.Content, &blog.PublishedAt, &blog.Archive, &blog.CommentCount); err != nil {
			return blogs, &m.ErrorMessage{
				Error:   "Row Scan Error",
				Details: []string{"Error scanning row into blog."},
				Code:    "",
			}
		}
	}

	if err := rows.Err(); err != nil {
		return blogs, &m.ErrorMessage{
			Error:   "Row Scan Error",
			Details: []string{"Error scanning row into blog."},
			Code:    "",
		}
	}
	return blogs, nil
}

func (db *BlogStore) Get(id string) (m.Blog, *m.ErrorMessage) {
	row := db.DB.QueryRow(`SELECT id, user_id, title, content, published_at, archive, comment_count FROM blogs WHERE id=?`, id)

	var blog m.Blog
	row.Scan(&blog.Id, &blog.UserID, &blog.Title, &blog.Content, &blog.PublishedAt, &blog.Archive, &blog.CommentCount)

	if err := row.Err(); err != nil {
		return m.Blog{}, &m.ErrorMessage{
			Error:   "Error scanning row.",
			Details: []string{"An Error occurred while scanning row."},
			Code:    "",
		}
	}
	return blog, nil
}

func (db *BlogStore) Post(blog m.Blog) *m.ErrorMessage {
	_, err := db.DB.Exec(`INSERT INTO blogs VALUES (?,?,?,?,?,?,?)`, blog.Id, blog.UserID, blog.Title, blog.Content, blog.PublishedAt, blog.Archive, blog.CommentCount)
	if err != nil {
		return &m.ErrorMessage{
			Error:   "Insertion Error",
			Details: []string{"Error inserting into the database!."},
			Code:    "",
		}
	}
	return nil
}

func (db *BlogStore) Patch(id string, body m.Blog) ([]string, *m.ErrorMessage) {
	//
	return nil, nil
}

func (db *BlogStore) Update(id string, blog m.Blog) *m.ErrorMessage {
	if _, err := db.DB.Exec(`UPDATE blogs SET title=?,content=?,published_at=?, comment_count=? WHERE id=?`, blog.Title, blog.Content, blog.PublishedAt, blog.CommentCount, id); err != nil {
		return &m.ErrorMessage{
			Error:   "Update insertion error",
			Details: []string{"An error occurred during insertion!."},
			Code:    "",
		}
	}
	return nil
}

func (db *BlogStore) Delete(id string) *m.ErrorMessage {
	if _, err := db.DB.Exec(`DELETE FROM blogs WHERE id=?`, id); err != nil {
		return &m.ErrorMessage{
			Error:   "Deletion error",
			Details: []string{"An error during deletion."},
			Code:    "",
		}
	}
	return nil
}

func (db *BlogStore) RegisterUser(user m.User) *m.ErrorMessage {
	return nil
}
func (db *BlogStore) LoginUser(email, password string) *m.ErrorMessage {
	return nil
}
