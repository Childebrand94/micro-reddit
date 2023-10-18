package utils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/models"
)

func CombinedPostComments(
	allPosts []models.Post,
	allComments []models.Comment,
) []models.PostWithComments {
	var result []models.PostWithComments

	for _, post := range allPosts {
		pwc := models.PostWithComments{
			Post: post,
		}
		for _, comment := range allComments {
			if comment.Post_ID == post.ID {
				pwc.Comments = append(pwc.Comments, comment)
			}
		}
		result = append(result, pwc)
	}

	return result
}

func GetVoteTotal(pool *pgxpool.Pool, postID int64) (pgtype.Int8, error) {
	var totalVotes pgtype.Int8
	query := `SELECT 
  SUM(CASE WHEN up_vote = true THEN 1 ELSE 0 END) - SUM(CASE WHEN up_vote = false THEN 1 ELSE 0 END)  as total_votes
	FROM post_votes 
	WHERE post_id = $1;`

	err := pool.QueryRow(context.TODO(), query, postID).Scan(&totalVotes)
	if err != nil {
		var zeroValue pgtype.Int8
		return zeroValue, err
	}

	fmt.Println(totalVotes)
	return totalVotes, nil
}
