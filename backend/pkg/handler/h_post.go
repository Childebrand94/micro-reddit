package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/database"
	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/Childebrand94/micro-reddit/pkg/utils"
)

type Post struct {
	DB *pgxpool.Pool
}

func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
	// use a model to decode the request into a struct
	var payload models.Post
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// decode request send error code if error
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to decode request", err)
		return
	}
	// hard code id until we have sessions
	id := 10

	// call database function to insert post into tables
	err = database.AddPostByUser(ctx, p.DB, int64(id), payload.URL, payload.Title)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to add user to database", err)
		return
	}

	// Send success response
	utils.SendSuccessfulResp(w, "Successfully created a Post")
}

func (p *Post) List(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	resp, err := database.GetAllPosts(ctx, p.DB)
	if err != nil {
		models.SendError(
			w,
			http.StatusInternalServerError,
			"Failed to fetch posts from database",
			err,
		)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to marshal data", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (p *Post) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	strID := chi.URLParam(r, "id")
	post_id := utils.ConvertID(strID, w)

	posts, comments, err := database.GetPostById(ctx, p.DB, int64(post_id))
	if err != nil {
		models.SendError(
			w,
			http.StatusInternalServerError,
			"Failed to get post data from database",
			err,
		)
		return
	}

	user, err := database.GetUserByID(ctx, p.DB, post_id)
	if err != nil {
		models.SendError(
			w,
			http.StatusInternalServerError,
			"Failed to get user data from database",
			err,
		)
		return
	}

	result := utils.ConstructPostResponses(posts, comments, user)
	data, err := json.Marshal(result[0])
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to prepare response", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (p *Post) PostVotes(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	user_id := 2
	strID := chi.URLParam(r, "id")
	post_id, err := strconv.Atoi(strID)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Failed to get id from URL", err)
	}
	strVote := chi.URLParam(r, "vote")

	var upVote bool

	if strVote == "up-vote" {
		upVote = true
	} else if strVote == "down-vote" {
		upVote = false
	} else {
		models.SendError(w, http.StatusBadRequest, "Bad URL", nil)
	}

	err = database.AddPostVotes(ctx, p.DB, int64(user_id), int64(post_id), upVote)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to get data from database", err)
		return
	}

	utils.SendSuccessfulResp(w, "Vote had been created")
}

func (p *Post) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a post by ID")
}

func (p *Post) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by ID")
}
