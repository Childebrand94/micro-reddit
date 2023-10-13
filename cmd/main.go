package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Childebrand94/micro-reddit/cmd/application"
	db "github.com/Childebrand94/micro-reddit/pkg/database"
	m "github.com/Childebrand94/micro-reddit/pkg/mock"
	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// Load in .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up database connection
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	// Starting Server
	fmt.Println("Starting Server....")
	app := application.New(pool)
	err = app.Start(context.TODO())
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func databaseTest(pool *pgxpool.Pool) {
	err := db.AddUser(pool, m.User1.First_name, m.User1.Last_name, m.User1.Email)
	if err != nil {
		fmt.Printf("Failed to execute AddUser: %v", err)
	}

	err = db.AddPostByUser(pool, m.User1.ID, m.Post1.URL, m.Post1.Title)
	if err != nil {
		fmt.Printf("Failed to execute AddPostByUser: %v", err)
	}

	var post *models.Post

	post, err = db.GetPostById(pool, m.Post1.ID)
	if err != nil {
		fmt.Printf("Failed to execute GetPost: %v", err)
	}

	fmt.Println("Author ID:", post.Author_ID)
	fmt.Println("Title:", post.Title)
	fmt.Println("URL:", post.URL)
}
