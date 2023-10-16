package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Comment struct {
	DB *pgxpool.Pool
}

func (c *Comment) Create(w http.ResponseWriter, r *http.Request) {
	// idStr := chi.URLParam(r, "id")
	// id, err := strconv.Atoi(idStr)
	// if err != nil {
	// 	models.SendError(w, http.StatusBadRequest, "Invalid ID", err)
	// 	return
	// }
	var payload models.Comment

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Bad request format", err)
		return
	}

}
