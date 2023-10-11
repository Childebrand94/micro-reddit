package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// User struct for demonstration
type User struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	DateJoined string `json:"date_joined"`
}

// App struct that holds the database pool
type App struct {
	Pool *pgxpool.Pool
}

// GetAllUsersHandler is a method on the App struct, so it can access the database pool.
func (app *App) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM users"
	rows, err := app.Pool.Query(context.Background(), query)
	if err != nil {
		http.Error(w, "Failed to query users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.DateJoined); err != nil {
			http.Error(w, "Failed to scan user", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	// Send users as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	// Load in .env file 
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	} 

	// Create the database pool (simplified, assumes DATABASE_URL is a valid connection string)
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	// Create the App with the database pool
	app := &App{Pool: pool}

	// Set up the HTTP server
	http.HandleFunc("/users", app.GetAllUsersHandler)
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}