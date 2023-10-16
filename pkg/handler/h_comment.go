package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Childebrand94/micro-reddit/pkg/database"
	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Comment struct {
	DB *pgxpool.Pool
}

func (c *Comment) Create(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
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

	err = database.AddComment(c.DB, int64(id), payload.Author_ID, payload.Parent_ID, payload.Message)
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
