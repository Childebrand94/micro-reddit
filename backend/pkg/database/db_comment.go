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
	query := `SELECT c.id, c.post_id, c.author_id, c.parent_id, c.message, 
                        c.created_at, u.first_name, u.last_name, u.username 
						FROM comments AS c  
						LEFT JOIN 
						users u ON u.id  = c.author_id 
						WHERE post_id = $1`

	rows, err := pool.Query(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.CommentResp

	for rows.Next() {
		var c models.CommentResp
		if err := rows.Scan(&c.ID, &c.Post_ID, &c.Author_ID, &c.Parent_ID, &c.Message, &c.Created_at, &c.Author.FirstName, &c.Author.LastName, &c.Author.UserName); err != nil {
			return nil, err
		}
		totalVotes, err := utils.GetVoteTotal(pool, c.ID, "comment_vote", "comment_id")
		if err != nil {
			return nil, err
		}
		c.Vote = totalVotes
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
	query := "INSERT INTO comment_vote (user_id, comment_id, up_vote) VALUES ($1, $2, $3)"
	_, err := pool.Exec(ctx, query, user_id, comment_id, up_vote)
	return err
}
