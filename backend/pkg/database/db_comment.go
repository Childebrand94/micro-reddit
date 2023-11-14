package database

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/Childebrand94/micro-reddit/pkg/utils"
)

func AddComment(ctx context.Context, pool *pgxpool.Pool, postID, authorID int64, parentID sql.NullInt64, message, path string) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}
	query := "INSERT INTO comments (post_id, author_id, parent_id, message, path) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	var id int64

	err = tx.QueryRow(ctx, query, postID, authorID, parentID, message, path).Scan(&id)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	if parentID.Valid {
		updateQuery := `UPDATE comments SET path = parent.path || '/' || CAST($2 AS TEXT) FROM comments AS parent WHERE parent.id = $1 AND comments.id = $2`
		_, err = tx.Exec(ctx, updateQuery, parentID.Int64, id)
	} else {
		updateQuery := `UPDATE comments SET path = CAST($1 AS TEXT) WHERE id = $1`
		_, err = tx.Exec(ctx, updateQuery, id)
	}
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return err
}

func ListComments(ctx context.Context, pool *pgxpool.Pool, postID int64, userId *int64) ([]models.CommentResp, error) {
	println("Listing Comments...")
	query := `SELECT c.id, c.post_id, c.author_id, c.parent_id, c.message, 
                        c.created_at, u.first_name, u.last_name, u.username 
						FROM comments AS c  
						LEFT JOIN 
						users u ON u.id  = c.author_id 
						WHERE post_id = $1
                        ORDER BY path`

	rows, err := pool.Query(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.CommentResp

	for rows.Next() {
		var c models.CommentResp
		if err := rows.Scan(
			&c.ID,
			&c.Post_ID,
			&c.Author_ID,
			&c.Parent_ID,
			&c.Message,
			&c.Path,
			&c.Created_at,
			&c.Author.FirstName,
			&c.Author.LastName,
			&c.Author.UserName,
		); err != nil {
			return nil, err
		}

		totalVotes, err := utils.GetVoteTotal(pool, c.ID, "s", "comment_id")
		if err != nil {
			return nil, err
		}

		c.Vote = totalVotes

		c.UsersVoteStatus, err = utils.UserCommentVoteCheck(ctx, pool, c.ID, userId)
		if err != nil {
			return nil, err
		}

		comments = append(comments, c)
	}
	return comments, nil
}

func AddCommentVotes(
	ctx context.Context,
	pool *pgxpool.Pool,
	user_id, comment_id int64,
	up_vote bool,
) error {
	query := `INSERT INTO comment_votes(user_id, comment_id, up_vote)
                VALUES ($1, $2, $3)
                ON CONFLICT (comment_id, user_id)
                DO UPDATE SET up_vote = EXCLUDED.up_vote;`
	_, err := pool.Exec(ctx, query, user_id, comment_id, up_vote)
	return err
}
