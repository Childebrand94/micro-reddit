package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/database"
	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/Childebrand94/micro-reddit/pkg/utils"
)

type Comment struct {
	DB *pgxpool.Pool
}

func (c *Comment) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	idStr := chi.URLParam(r, "id")
	post_id, err := strconv.Atoi(idStr)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	var payload models.Comment

	//    var payload models.Comment
	// parent_id := sql.NullInt64{
	// 	Int64: payload.Parent_ID,
	// 	Valid: true,
	// }

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to decode payload", err)
		return
	}

	err = database.AddComment(
		ctx,
		c.DB,
		int64(post_id),
		payload.Author_ID,
		sql.NullInt64(payload.Parent_ID),
		payload.Message,
	)
	if err != nil {
		models.SendError(
			w,
			http.StatusInternalServerError,
			"Failed to add comments to database",
			err,
		)
		return
	}

	utils.SendSuccessfulResp(w, "Comment Created")
}

func (c *Comment) List(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	idStr := chi.URLParam(r, "id")
	post_id := utils.ConvertID(idStr, w)

	comments, err := database.ListComments(ctx, c.DB, int64(post_id))
	if err != nil {
		models.SendError(
			w,
			http.StatusInternalServerError,
			"Failed to get comments form database",
			err,
		)
		return
	}
	data, err := json.Marshal(comments)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to marshal data", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (c *Comment) CommentVotes(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	commentID_str := chi.URLParam(r, "comment_id")
	vote_param := chi.URLParam(r, "vote")

	// userID hard coded until sessions
	user_id := 2
	// convert id to int
	comment_id := utils.ConvertID(commentID_str, w)

	var vote bool
	if vote_param == "up-vote" {
		vote = true
	} else {
		vote = false
	}

	err := database.AddCommentVotes(ctx, c.DB, int64(user_id), int64(comment_id), vote)
	if err != nil {
		models.SendError(
			w,
			http.StatusInternalServerError,
			"Failed to get comments form database",
			err,
		)
		return
	}

	utils.SendSuccessfulResp(w, "Votes on Comments successful")
}
