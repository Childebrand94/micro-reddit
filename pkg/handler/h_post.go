package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/database"
	"github.com/Childebrand94/micro-reddit/pkg/models"
)

type Post struct {
	DB *pgxpool.Pool
}

func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
	// use a model to decode the request into a struct
	var payload models.Post

	// decode request send error code if error
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to decode request", err)
		return
	}
	// hard code id until we have sessions
	id := 2

	// call database function to insert post into tables
	err = database.AddPostByUser(p.DB, int64(id), payload.URL)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to add user to database", err)
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Post created successfully",
	})
}

func (p *Post) List(w http.ResponseWriter, r *http.Request) {
	allComments, allPosts, err := database.GetAllPosts(p.DB)
	if err != nil {
		models.SendError(
			w,
			http.StatusInternalServerError,
			"Failed to fetch posts from database",
			err,
		)
		return
	}

	type postWithComments struct {
		models.Post
		Comments []models.Comment `json:"Comments"`
	}

	var result []postWithComments

	for _, post := range allPosts {
		pwc := postWithComments{
			Post: post,
		}
		for _, comment := range allComments {
			if comment.Post_ID == post.ID {
				pwc.Comments = append(pwc.Comments, comment)
			}
		}
		result = append(result, pwc)
	}

	data, err := json.Marshal(result)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to marshal data", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (p *Post) GetByID(w http.ResponseWriter, r *http.Request) {
	strID := chi.URLParam(r, "id")
	post_id, err := strconv.Atoi(strID)
	println(post_id)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Failed to get id from URL", err)
	}

	post, err := database.GetPostById(p.DB, int64(post_id))
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to get data from database", err)
	}

	data, err := json.Marshal(post)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to prepare response", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (p *Post) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a post by ID")
}

func (p *Post) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by ID")
}
