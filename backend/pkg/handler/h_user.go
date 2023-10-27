package handler

import (
	"context"
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

type User struct {
	DB *pgxpool.Pool
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*3)
	defer ctxCancel()

	var payload models.User
	// Decode the request body to get user details
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Bad request format", err)
		return
	}
	defer r.Body.Close()

	err = database.AddUser(ctx, u.DB, payload)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to add user to database", err)
		return
	}

	utils.SendSuccessfulResp(w, "Successfully created a user")
}

func (u *User) List(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*3)
	defer ctxCancel()
	users, err := database.GetAllUsers(ctx, u.DB)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed get users form database", err)
		return
	}

	data, err := json.Marshal(users)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to marshal data", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (u *User) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*3)
	defer ctxCancel()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	user, err := database.GetUserWithCPByID(ctx, u.DB, id)
	if err != nil {
		models.SendError(
			w,
			http.StatusInternalServerError,
			"Unable to fetch user form database",
			err,
		)
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to process user data", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (u *User) UpdateByID(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*3)
	defer ctxCancel()

	idStr := chi.URLParam(r, "id")
	id := utils.ConvertID(idStr, w)
	var updateUser models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateUser)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Bad Request", err)
		return
	}
	err = database.UpdateUserByID(ctx, u.DB, updateUser, int64(id))
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to update database", err)
		return
	}

	utils.SendSuccessfulResp(w, "User was successfully updated")
}