package database

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/Childebrand94/micro-reddit/pkg/utils"
)

func AddComment(
	ctx context.Context,
	pool *pgxpool.Pool,
	postID, authorID int64,
	parentID sql.NullInt64,
	message string,
) error {
	_, err := pool.Exec(ctx,
		`Insert INTO comments (post_id, author_id, parent_id, message)
					Values($1, $2, $3, $4)`, postID, authorID, parentID, message)
	return err
}

func ListComments(ctx context.Context, pool *pgxpool.Pool, postID int64) ([]models.CommentResp, error) {
	query := `SELECT * FROM comments WHERE post_id = $1`
	rows, err := pool.Query(ctx, query, postID)
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
		totalVotes, err := utils.GetVoteTotal(pool, c.ID, "comment_vote", "comment_id")
		if err != nil {
			return nil, err
		}
		c.Vote = totalVotes
		comments = append(comments, c)
	}

	users, err := GetAllUsers(ctx, pool)
	if err != nil {
		return nil, err
	}

	commentResp := utils.AddAuthorComment(comments, users)

	return commentResp, nil
}

func AddCommentVotes(
	ctx context.Context,
	pool *pgxpool.Pool,
	user_id, comment_id int64,
	up_vote bool,
) error {
	query := "INSERT INTO comment_vote (user_id, comment_id, up_vote) VALUES ($1, $2, $3)"
	_, err := pool.Exec(ctx, query, user_id, comment_id, up_vote)
	return err
}
