package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/database"
	"github.com/Childebrand94/micro-reddit/pkg/models"
)

type Comment struct {
	DB *pgxpool.Pool
}

func (c *Comment) Create(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	post_id, err := strconv.Atoi(idStr)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Invalid ID", err)
		return
	}
	var payload models.Comment

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Bad request format", err)
		return
	}

	err = database.AddComment(c.DB, int64(post_id), payload.Author_ID, payload.Parent_ID, payload.Message)
	if err != nil {
		models.SendError(
			w,
			http.StatusInternalServerError,
			"Failed to add comments to database",
			err,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Comment created successfully",
	})
}

func (c *Comment) List(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	post_id, err := strconv.Atoi(idStr)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	comments, err := database.ListComments(c.DB, int64(post_id))
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
