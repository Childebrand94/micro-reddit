package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Childebrand94/micro-reddit/pkg/database"
	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/jackc/pgx/v5/pgxpool"
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
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// hard code id until we have sessions
	id := 1

	// call database function to insert post into tables
	err = database.AddPostByUser(p.DB, int64(id), payload.Title, payload.URL)
	if err != nil {
		fmt.Printf("Failed to add user to database %v", err)
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Post created successfully",
	})
}

func (p *Post) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all posts")
}

func (p *Post) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get post by ID")
}

func (p *Post) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a post by ID")
}

func (p *Post) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by ID")
}
