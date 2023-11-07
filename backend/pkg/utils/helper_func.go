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
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/models"
)

func GetVoteTotal(pool *pgxpool.Pool, id int64, table, column string) (int, error) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*3)
	defer ctxCancel()

	var totalVotes int
	query := fmt.Sprintf(`SELECT 
  COALESCE(SUM(CASE WHEN up_vote = true THEN 1 ELSE 0 END) - SUM(CASE WHEN up_vote = false THEN 1 ELSE 0 END), 0) as total_votes
	FROM %s 
	WHERE %s = $1;`, table, column)

	err := pool.QueryRow(ctx, query, id).Scan(&totalVotes)
	if err != nil {
		return 0, err
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

func GenerateSessionToken() string {
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

func GetSessionCookie(r *http.Request) (*http.Cookie, *models.CustomError) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// No session Cookie present
			return nil, &models.CustomError{
				StatusCode:    http.StatusUnauthorized,
				Message:       "Session cookie not found",
				OriginalError: err,
			}
		}
		return nil, &models.CustomError{
			StatusCode:    http.StatusInternalServerError,
			Message:       "Error retrieving session cookie",
			OriginalError: err,
		}
	}

	return cookie, nil
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

func CalcKarma(posts []models.PostWithAuthor, comments []models.CommentResp) int {
	var totalPostVotes int
	var totalCommentVotes int
	for _, p := range posts {
		totalPostVotes = 0
		totalPostVotes += p.Vote
	}

	for _, c := range comments {
		totalCommentVotes = 0
		totalCommentVotes += c.Vote
	}

	karma := totalPostVotes + totalCommentVotes + len(posts) + len(comments)
	return karma
}

type SortMethod struct {
	ByVotesDesc    string
	ByCreationDesc string
	ByHot          string
}

func GetSortMethod(sortType string) string {
	baseQuery := `
        SELECT 
            p.id,
            p.author_id,
            p.title,
            p.url,
            p.created_at,
            p.updated_at,
            u.first_name,
            u.last_name,
            u.username,
            COALESCE(SUM(CASE WHEN pv.up_vote THEN 1 ELSE 0 END), 0) AS upvotes_count
        FROM 
            posts p
        LEFT JOIN 
            users u ON u.id = p.author_id
        LEFT JOIN 
            post_votes pv ON p.id = pv.post_id
        GROUP BY 
            p.id, u.id
    `

	sortMethods := SortMethod{
		ByVotesDesc:    baseQuery + " ORDER BY upvotes_count DESC, p.created_at DESC",
		ByCreationDesc: baseQuery + " ORDER BY p.created_at DESC",
		ByHot:          baseQuery,
	}

	switch sortType {
	case "top":
		return sortMethods.ByVotesDesc
	case "new":
		return sortMethods.ByCreationDesc
	case "hot":
		return sortMethods.ByHot

	default:
		return sortMethods.ByCreationDesc
	}
}

//
// func URLFormatter(str string) (*string, error) {
//     u, err := url.Parse(str)
//     if err != nil {
//         return nil, err
//     }
//     if u.Scheme == ""{
//         u.Scheme = "http"
//     }
//     u.Q
//
// }
