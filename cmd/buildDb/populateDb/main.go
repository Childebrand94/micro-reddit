package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/Childebrand94/micro-reddit/pkg/models"
)

func main() {
	// Load in .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up database connection
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	users := populateDatabaseWithUsers(dbpool)
	populateDatabaseWithPosts(dbpool, users)
}

func populateDatabaseWithUsers(pool *pgxpool.Pool) []models.User {
	var users []models.User

	for i := 0; i < 10; i++ {
		var u models.User

		u.First_name = fmt.Sprintf("user%d", i)
		u.Last_name = fmt.Sprintf("user%d last_name", i)
		u.Email = fmt.Sprintf("user%d.com", i)
		u.Username = fmt.Sprintf("username%d", i)

		users = append(users, u)
	}
	batch := &pgx.Batch{}

	for _, u := range users {
		batch.Queue(
			"INSERT INTO users (first_name, last_name, username, email) VALUES ($1, $2, $3, $4)",
			u.First_name,
			u.Last_name,
			u.Username,
			u.Email,
		)
	}

	br := pool.SendBatch(context.TODO(), batch)
	_, err := br.Exec()
	if err != nil {
		log.Fatalf("Failed to add users to database: %v", err)
	}

	return users
}

func populateDatabaseWithPosts(pool *pgxpool.Pool, users []models.User) {
	var posts []models.Post

	for _, u := range users {
		var p models.Post

		p.Author_ID = u.ID
		p.URL = fmt.Sprintf("www.postByUser%d.com", u.ID)

		posts = append(posts, p)
	}
	batch := &pgx.Batch{}

	for _, p := range posts {
		batch.Queue("INSERT INTO posts (authour_id, url) VALUES ($1, $2)", p.Author_ID, p.URL)
	}

	br := pool.SendBatch(context.TODO(), batch)
	_, err := br.Exec()
	if err != nil {
		log.Fatalf("Failed to add posts to database: %v", err)
	}
}
