package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func AddComment(pool *pgxpool.Pool, postID, authorID int64, parentID *int64, message string) error {
	_, err := pool.Exec(context.TODO(),
		`Insert INTO comments (post_id, author_id, parent_id, message)
					Values($1, $2, $3, $4)`, postID, authorID, parentID, message)
	return err
}
