package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/Childebrand94/micro-reddit/pkg/utils"
)

type Session struct {
	DB *pgxpool.Pool
}

func (s *Session) IsSession(w http.ResponseWriter, r *http.Request) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*3)
	defer ctxCancel()

	// Check for cookie if no cookie is found return 200 with loggedIn set to false
	cookie, customErr := utils.GetSessionCookie(r)
	if customErr != nil {
		response := map[string]interface{}{
			"loggedIn": false,
		}
		if customErr.OriginalError == http.ErrNoCookie {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		} else {
			models.SendError(w, customErr.StatusCode, customErr.Message, customErr.OriginalError)
		}
		return
	}

	session, err := utils.ValidateSessionToken(ctx, s.DB, cookie.Value)
	if err != nil {
		println("reset cookie")
		// if there is no session but a cookie is found clear the cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    "",
			Path:     "/",
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
			SameSite: http.SameSiteNoneMode,
		})
		models.SendError(w, http.StatusNotFound, "Session not found", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"loggedIn": true,
		"userId":   session.Session_id,
	}
	json.NewEncoder(w).Encode(response)
	return
}
