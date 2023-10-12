package handler

import (
	"fmt"
	"net/http"
  "encoding/json"
  
  "github.com/Childebrand94/micro-reddit/pkg/database"
	"github.com/jackc/pgx/v5/pgxpool"
  "github.com/Childebrand94/micro-reddit/pkg/models"
)

type User struct{
  DB *pgxpool.Pool
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
  var payload models.User
  	// Decode the request body to get user details
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

 err = database.AddUser(u.DB, payload.First_name, payload.Last_name, payload.Email)
  if err != nil {
   fmt.Printf("Failed to add user to database %v", err) 
  }

	// Send a success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
	})
}

func (u *User) List(w http.ResponseWriter, r *http.Request) {
  fmt.Println("List all posts")
}

func (u *User) GetByID(w http.ResponseWriter, r *http.Request){
  fmt.Println("Get post by ID")
}

func (u *User) UpdateByID(w http.ResponseWriter, r *http.Request){
  fmt.Println("Update a post by ID")
}

func (u *User) DeleteByID(w http.ResponseWriter, r *http.Request){
  fmt.Println("Delete an order by ID")
}

