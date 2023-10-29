package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"

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

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Bad request format", err)
		return
	}
	defer r.Body.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(payload.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to process user password", err)
		return
	}

	payload.Password = string(hashedPassword)

	err = database.AddUser(ctx, u.DB, payload)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to add user to database", err)
		return
	}

	log.Print("Successfully created user")

	// Crate Session for user
	user, err := database.GetUserByEmail(ctx, u.DB, payload.Email)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to get user by email", err)
		return
	}

	sessionId := utils.GenereateSessionToken()

	err = database.CreateSession(ctx, u.DB, sessionId, user.ID)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to generate session", err)
		return
	}

	utils.SetSessionToken(w, sessionId)

	utils.SendSuccessfulResp(w, "Successfully created a user and session")
}

func (u *User) Authenticate(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*3)
	defer ctxCancel()

	type LoginRequest struct {
		Email    string
		Password string
	}

	var payload LoginRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Bad request format", err)
		return
	}

	user, err := database.GetUserByEmail(ctx, u.DB, payload.Email)
	if err != nil {
		models.SendError(w, http.StatusUnauthorized, "Email not registered", err)
		return
	}

	hashedPassword := user.Password
	enteredPassword := payload.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(enteredPassword))
	if err != nil {
		models.SendError(w, http.StatusUnauthorized, "Invalid credentials", err)
		return
	}

	sessionId := utils.GenereateSessionToken()

	err = database.CreateSession(ctx, u.DB, sessionId, user.ID)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to generate session", err)
		return
	}

	utils.SetSessionToken(w, sessionId)

	utils.SendSuccessfulResp(w, "Successfully created session")
}

func (u *User) Logout(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*3)
	defer ctxCancel()

	cookie, err := r.Cookie("session_token")
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "No active session found", err)
		return
	}
	sessionId := cookie.Value

	err = database.DeleteSession(ctx, u.DB, sessionId)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to delete session", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})

	utils.SendSuccessfulResp(w, "Successfully logged out")
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
