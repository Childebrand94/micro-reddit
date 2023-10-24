package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/models"
)

func ConstructPostResponses(allPosts []models.Post, allComments []models.Comment, allUsers []models.User) []models.PostResponse {
	var result []models.PostResponse

	for _, post := range allPosts {
		pr := models.PostResponse{
			Post: post,
		}
		for _, comment := range allComments {
			if comment.Post_ID == post.ID {
				pr.Comments = append(pr.Comments, comment)
			}
		}
		for _, user := range allUsers {
			if post.Author_ID == user.ID {
				pr.Author.FirstName = user.First_name
				pr.Author.LastName = user.Last_name
				pr.Author.UserName = user.Username
			}
		}
		result = append(result, pr)
	}

	return result
}

func GetVoteTotal(pool *pgxpool.Pool, id int64, table, column string) (pgtype.Int8, error) {
	var totalVotes pgtype.Int8
	query := fmt.Sprintf(`SELECT 
  SUM(CASE WHEN up_vote = true THEN 1 ELSE 0 END) - SUM(CASE WHEN up_vote = false THEN 1 ELSE 0 END)  as total_votes
	FROM %s 
	WHERE %s = $1;`, table, column)

	err := pool.QueryRow(context.TODO(), query, id).Scan(&totalVotes)
	if err != nil {
		var zeroValue pgtype.Int8
		return zeroValue, err
	}

	return totalVotes, nil
}

func ConvertID(s string, w http.ResponseWriter) int {
	intID, err := strconv.Atoi(s)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Invalid ID", err)
		os.Exit(1)
	}
	return intID
}

func SendSuccessfulResp(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": message,
	})
}
