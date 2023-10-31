package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/models"
)

// func ConstructPostResponses(
//
//	allPosts []models.Post,
//	allComments []models.CommentResp,
//	allUsers []models.User,
//
//	) []models.PostResponse {
//		var result []models.PostResponse
//
//		for _, post := range allPosts {
//			pr := models.PostResponse{
//				Post: post,
//			}
//			for _, comment := range allComments {
//				if comment.Post_ID == post.ID {
//					pr.Comments = append(pr.Comments, comment)
//				}
//			}
//			for _, user := range allUsers {
//				if post.Author_ID == user.ID {
//					pr.Author.FirstName = user.First_name
//					pr.Author.LastName = user.Last_name
//					pr.Author.UserName = user.Username
//				}
//			}
//			result = append(result, pr)
//		}
//
//		return result
//	}
//
//	func AddAuthorComment(allComments []models.Comment, allUsers []models.User) []models.CommentResp {
//		var result []models.CommentResp
//
//		for _, comment := range allComments {
//			cr := models.CommentResp{
//				Comment: comment,
//			}
//			for _, user := range allUsers {
//				if comment.Author_ID == user.ID {
//					cr.Author.FirstName = user.First_name
//					cr.Author.LastName = user.Last_name
//					cr.Author.UserName = user.Username
//				}
//			}
//			result = append(result, cr)
//		}
//		return result
//	}
//
//	func AddAuthorPosts(allPosts []models.Post, user models.User) []models.PostWithAuthor {
//		var result []models.PostWithAuthor
//
//		for _, post := range allPosts {
//			pr := models.PostWithAuthor{
//				Post: post,
//			}
//			pr.Author.FirstName = user.First_name
//			pr.Author.LastName = user.Last_name
//			pr.Author.UserName = user.Username
//
//			result = append(result, pr)
//		}
//		return result
//	}
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

func GenereateSessionToken() string {
	token := uuid.New().String()
	return token
}

func SetSessionToken(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(12 * time.Hour),
		SameSite: http.SameSiteNoneMode,
	})
}

func GetSessionCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// No session Cookie present
			models.SendError(w, http.StatusUnauthorized, "Not authorized", err)
			return nil
		}
		models.SendError(w, http.StatusInternalServerError, "Internal server error", err)
		return nil
	}

	return cookie
}

func ValidateSessionToken(ctx context.Context, pool *pgxpool.Pool, token string) (*models.Session, error) {
	query := "SELECT session_id, user_id FROM sessions WHERE session_id = $1"

	var s models.Session

	row := pool.QueryRow(ctx, query, token)
	err := row.Scan(&s.Session_id, &s.User_id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("session not found")
		}
		return nil, err
	}

	return &s, err
}
