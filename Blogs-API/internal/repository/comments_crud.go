package repository

import (
	m "blog/pkg/models"
	h "blog/pkg/utils"
)

/*
ListComments() ([]m.Comment, *m.ErrorMessage)
GetComment(id string) (m.Comment, *m.ErrorMessage)
	PostComment(id string) (m.Comment, *m.ErrorMessage)
	PatchComment(id string, query map[string]string) (m.Comment, *m.ErrorMessage)
	UpdateComment(id string, comment m.Comment) *m.ErrorMessage
	DeleteComment(id string) *m.ErrorMessage
*/

func (db *BlogStore) ListComments(blog_id string) ([]m.Comment, *m.ErrorMessage) {
	rows, err := db.DB.Query(h.LIST_COMMENT_QUERY, blog_id)
	if err != nil {
		return nil, &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{h.SERVER_ERROR_DETAIL},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	defer rows.Close()

	var comments []m.Comment

	for rows.Next() {
		var comment m.Comment
		err := rows.Scan(&comment.Id, &comment.Content, &comment.CreatedAt)
		if err != nil {
			
		}
		comments = append(comments, comment)
	}

	return nil, nil
}

func (db *BlogStore) GetComment(blog_id, id string) (m.Comment, *m.ErrorMessage) {
	return m.Comment{}, nil
}

func (db *BlogStore) PostComment(blog_id, id string) *m.ErrorMessage {
	return nil
}

func (db *BlogStore) PatchComment(blog_id, id string, query map[string]string) *m.ErrorMessage {
	return nil
}

func (db *BlogStore) UpdateComment(blog_id, id string, comment m.Comment) *m.ErrorMessage {
	return nil
}

func (db *BlogStore) DeleteComment(blog_id, id string) *m.ErrorMessage {
	return nil
}
