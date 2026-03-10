package repository

import (
	m "blog/pkg/models"
	h "blog/pkg/utils"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func (db *BlogStore) ListComments(ctx context.Context, blog_id string) ([]m.Comment, *m.ErrorMessage) {
	rows, err := db.DB.QueryContext(ctx, h.LIST_COMMENTS_QUERY, blog_id)
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
		err2 := rows.Scan(&comment.Id, &comment.Content, &comment.CreatedAt)
		if err2 != nil {
			if errors.Is(err2, sql.ErrConnDone) {
				return nil, &m.ErrorMessage{
					Error:   h.CONN_ERR,
					Details: []string{err2.Error()},
					Code:    h.SERVER_ERROR_CODE,
				}
			}
			return nil, &m.ErrorMessage{
				Error:   h.SERVER_ERROR,
				Details: []string{err2.Error()},
				Code:    h.SERVER_ERROR_CODE,
			}
		}
		comments = append(comments, comment)
	}

	if err3 := rows.Err(); err3 != nil {
		return nil, &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{err3.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	return comments, nil
}

func (db *BlogStore) GetComment(ctx context.Context, blog_id, id string) (m.Comment, *m.ErrorMessage) {
	row := db.DB.QueryRowContext(ctx, h.GET_COMMENT_QUERY, blog_id, id)

	var comment m.Comment

	if err := row.Scan(&comment.Id, &comment.BlogID, &comment.Content, &comment.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrConnDone) {
			return m.Comment{}, &m.ErrorMessage{
				Error:   h.CONN_ERR,
				Details: []string{err.Error()},
				Code:    h.SERVER_ERROR_CODE,
			}
		} else if errors.Is(err, sql.ErrNoRows) {
			return m.Comment{}, &m.ErrorMessage{
				Error:   h.NOT_FOUND_ERR,
				Details: []string{h.NOT_FOUND_ERR_DETAIL},
				Code:    h.NOT_FOUND_ERR_CODE,
			}
		} else {
			return m.Comment{}, &m.ErrorMessage{
				Error:   h.SERVER_ERROR,
				Details: []string{h.SERVER_ERROR_DETAIL},
				Code:    h.SERVER_ERROR_CODE,
			}
		}
	}

	if err2 := row.Err(); err2 != nil {
		return m.Comment{}, &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{err2.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	return comment, nil
}

func (db *BlogStore) PostComment(ctx context.Context, blog_id, id string, comment m.Comment) *m.ErrorMessage {
	result, err := db.DB.ExecContext(ctx, h.POST_COMMENT_QUERY, comment.Id, comment.BlogID, comment.Content)
	if err != nil {
		if errors.Is(err, sql.ErrConnDone) {
			return &m.ErrorMessage{
				Error:   h.CONN_ERR,
				Details: []string{err.Error()},
				Code:    h.SERVER_ERROR_CODE,
			}
		}
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{err.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	rowsAffected, err2 := result.RowsAffected()
	if err2 != nil {
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
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

	href := fmt.Sprintf("blogs/%s/comments/%s", comment.BlogID, comment.Id)
	link := m.Links{
		Id:             uuid.NewString(),
		BlogID:         comment.BlogID,
		Relationship:   "self",
		HyperReference: href,
	}

	linkResult, err := db.DB.ExecContext(ctx, h.POST_COMMENT_LINK_QUERY, link.Id, link.BlogID, link.Relationship, link.HyperReference)
	if err != nil {
		if errors.Is(err, sql.ErrConnDone) {
			return &m.ErrorMessage{
				Error:   h.CONN_ERR,
				Details: []string{err.Error()},
				Code:    h.SERVER_ERROR_CODE,
			}
		}
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{err.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	rowsAffected2, err3 := linkResult.RowsAffected()
	if err3 != nil {
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{err3.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	if rowsAffected2 == 0 {
		return &m.ErrorMessage{
			Error:   h.NOT_FOUND_ERR,
			Details: []string{h.NOT_FOUND_ERR_DETAIL},
			Code:    h.NOT_FOUND_ERR_CODE,
		}
	}

	err4 := db.IncrementCommentCount(ctx, blog_id)
	if err4 != nil {
		if err4.Error == h.CONN_ERR {
			return &m.ErrorMessage{
				Error:   h.CONN_ERR,
				Details: []string{err.Error()},
				Code:    h.SERVER_ERROR_CODE,
			}
		}
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{err.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	return nil
}

func (db *BlogStore) PatchComment(ctx context.Context, blog_id, id string, query map[string]string) *m.ErrorMessage {
	comment, err := db.GetComment(ctx, blog_id, id)
	if err != nil {
		return err
	}

	key := "content"
	if _, ok := query[key]; !ok {
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{h.SERVER_ERROR_DETAIL},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	comment.Content = query[key]

	err2 := db.UpdateComment(ctx, blog_id, id, comment)
	if err2 != nil {
		return err2
	}

	return nil
}

func (db *BlogStore) UpdateComment(ctx context.Context, blog_id, id string, comment m.Comment) *m.ErrorMessage {
	result, err := db.DB.ExecContext(ctx, h.UPDATE_COMMENT_QUERY, comment.Content, blog_id, id)
	if err != nil {
		if errors.Is(err, sql.ErrConnDone) {
			return &m.ErrorMessage{
				Error:   h.CONN_ERR,
				Details: []string{err.Error()},
				Code:    h.SERVER_ERROR_CODE,
			}
		}
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{err.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	rowsAffected, err2 := result.RowsAffected()
	if err2 != nil {
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{h.SERVER_ERROR_DETAIL},
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

func (db *BlogStore) DeleteComment(ctx context.Context, blog_id, id string) *m.ErrorMessage {
	result, err := db.DB.ExecContext(ctx, h.DELETE_COMMENT_QUERY, blog_id, id)
	if err != nil {
		if errors.Is(err, sql.ErrConnDone) {
			return &m.ErrorMessage{
				Error:   h.CONN_ERR,
				Details: []string{err.Error()},
				Code:    h.SERVER_ERROR_CODE,
			}
		}
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{err.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	rowsAffected, err2 := result.RowsAffected()
	if err2 != nil {
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{h.SERVER_ERROR_DETAIL},
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

	err3 := db.DecrementCommentCount(ctx, blog_id)
	if err3 != nil {
		if err3.Error == h.CONN_ERR {
			return &m.ErrorMessage{
				Error:   h.CONN_ERR,
				Details: []string{err.Error()},
				Code:    h.SERVER_ERROR_CODE,
			}
		}
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{err.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}
	return nil
}

func (db *BlogStore) IncrementCommentCount(ctx context.Context, blog_id string) *m.ErrorMessage {
	result, err := db.DB.ExecContext(ctx, h.UPDATE_COMMENT_COUNT_QUERY_INCR, blog_id)
	if err != nil {
		if errors.Is(err, sql.ErrConnDone) {
			return &m.ErrorMessage{
				Error:   h.CONN_ERR,
				Details: []string{err.Error()},
				Code:    h.SERVER_ERROR_CODE,
			}
		}
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{err.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	rowsAffected, err2 := result.RowsAffected()
	if err2 != nil {
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{h.SERVER_ERROR_DETAIL},
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

func (db *BlogStore) DecrementCommentCount(ctx context.Context, blog_id string) *m.ErrorMessage {
	result, err := db.DB.ExecContext(ctx, h.UPDATE_COMMENT_COUNT_QUERY_DESC, blog_id)
	if err != nil {
		if errors.Is(err, sql.ErrConnDone) {
			return &m.ErrorMessage{
				Error:   h.CONN_ERR,
				Details: []string{err.Error()},
				Code:    h.SERVER_ERROR_CODE,
			}
		}
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{err.Error()},
			Code:    h.SERVER_ERROR_CODE,
		}
	}

	rowsAffected, err2 := result.RowsAffected()
	if err2 != nil {
		return &m.ErrorMessage{
			Error:   h.SERVER_ERROR,
			Details: []string{h.SERVER_ERROR_DETAIL},
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
