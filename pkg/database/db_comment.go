package database

import (
	"context"

	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AddComment(pool *pgxpool.Pool, postID, authorID, parentID pgtype.Int8, message pgtype.Text) error {
	_, err := pool.Exec(context.TODO(),
		`Insert INTO comments (post_id, author_id, parent_id, message)
					Values($1, $2, $3, $4)`, postID, authorID, parentID, message)
	return err
}

func ListComments(pool *pgxpool.Pool, postID pgtype.Int8) ([]models.Comment, error) {
	query := `SELECT * FROM comments WHERE post_id = $1`
	rows, err := pool.Query(context.TODO(), query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment

	for rows.Next() {
		var c models.Comment
		if err := rows.Scan(&c.ID, &c.Post_ID, &c.Author_ID, &c.Parent_ID, &c.Message, &c.Created_at); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}
