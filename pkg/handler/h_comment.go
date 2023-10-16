package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Childebrand94/micro-reddit/pkg/database"
	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Comment struct {
	DB *pgxpool.Pool
}

func (c *Comment) Create(w http.ResponseWriter, r *http.Request) {
	var post_id pgtype.Int8
	idStr := chi.URLParam(r, "id")
	err := post_id.Scan(idStr)
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

	err = database.AddComment(c.DB, post_id, payload.Author_ID, payload.Parent_ID, payload.Message)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to add comments to database", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Comment created successfully",
	})

}

func (c *Comment) List(w http.ResponseWriter, r *http.Request) {
	var post_id pgtype.Int8
	idStr := chi.URLParam(r, "id")
	err := post_id.Scan(idStr)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	comments, err := database.ListComments(c.DB, post_id)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to get comments form database", err)
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
