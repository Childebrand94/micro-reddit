package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Childebrand94/micro-reddit/pkg/database"
	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
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
	users, err := database.GetAllUsers(u.DB)
	if err != nil {
		http.Error(w, "Failed to Get Users", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Failed to match", err)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	w.Write(data)

}

func (u *User) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id , err := strconv.Atoi(idStr)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := database.GetUserByID(u.DB, id)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Unable to fetch user form database")
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to process user data")
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

// func (u *User) UpdateByID(w http.ResponseWriter, r *http.Request) {
// 	idStr := chi.URLParam(r, "id")
// 	id , err := strconv.Atoi(idStr)
// 	if err != nil {
// 		models.SendError(w, http.StatusBadRequest, "Invalid user ID")
// 		return
// 	}


// }

func (u *User) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by ID")
}
