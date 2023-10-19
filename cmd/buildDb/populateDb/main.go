package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/Childebrand94/micro-reddit/pkg/database"
	"github.com/Childebrand94/micro-reddit/pkg/models"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

	populateDatabaseWithUsers(dbpool)
	populateDatabaseWithPosts(dbpool)
	addVotesPosts(dbpool)
	populateDatabaseWithComments(dbpool)
	populateCommentsWithComments(dbpool)
	fmt.Println("Successfully populated database.")
}

func populateDatabaseWithUsers(pool *pgxpool.Pool) {
	var users []models.User

	for i := 0; i < 10; i++ {
		var u models.User

		u.First_name = faker.FirstName()
		u.Last_name = faker.LastName()
		u.Email = faker.Email()
		u.Username = faker.Username()

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
	// fmt.Printf("Users: %v", users)
}

func populateDatabaseWithPosts(pool *pgxpool.Pool) {
	users, err := database.GetAllUsers(pool)
	if err != nil {
		log.Fatalf("Failed to get users from database: %v", err)
	}
	var posts []models.Post

	for _, u := range users {
		var p models.Post

		p.Author_ID = u.ID
		p.URL = faker.URL()

		posts = append(posts, p)
	}
	batch := &pgx.Batch{}

	for _, p := range posts {
		batch.Queue("INSERT INTO posts (author_id, url) VALUES ($1, $2)", p.Author_ID, p.URL)
	}

	br := pool.SendBatch(context.TODO(), batch)
	_, err = br.Exec()
	if err != nil {
		log.Fatalf("Failed to add posts to database: %v", err)
	}
}

func addVotesPosts(pool *pgxpool.Pool) {
	posts, err := database.GetPostsHelper(pool)
	if err != nil {
		log.Fatalf("Failed to get posts from database: %v", err)
	}
	users, err := database.GetAllUsers(pool)
	if err != nil {
		log.Fatalf("Failed to get users from database: %v", err)
	}
	type vote struct {
		User_id int64
		Post_id int64
		Up_vote bool
	}

	var votes []vote

	for _, p := range posts {
		for _, u := range users {
			var v vote

			v.User_id = u.ID
			v.Post_id = p.ID
			v.Up_vote = randomBool()

			votes = append(votes, v)
		}
	}
	for _, v := range votes {
		database.AddPostVotes(pool, v.User_id, v.Post_id, v.Up_vote)
	}

	batch := &pgx.Batch{}

	for _, v := range votes {
		batch.Queue(
			"INSERT INTO post_votes (user_id, post_id, up_vote) VALUES ($1, $2, $3)",
			v.User_id,
			v.Post_id,
			v.Up_vote,
		)
	}

	br := pool.SendBatch(context.TODO(), batch)
	_, err = br.Exec()
	if err != nil {
		log.Fatalf("Failed to add votes to database: %v", err)
	}
}

func populateDatabaseWithComments(pool *pgxpool.Pool) {
	posts, err := database.GetPostsHelper(pool)
	if err != nil {
		log.Fatalf("Failed to get posts from database: %v", err)
	}
	users, err := database.GetAllUsers(pool)
	if err != nil {
		log.Fatalf("Failed to get users from database: %v", err)
	}
	// query := "ALTER TABLE comments ALTER COLUMN parent_id DROP NOT NULL;"

	// _, err = pool.Exec(context.TODO(), query)
	// if err != nil {
	// 	log.Fatalf("Faild to remove constraint %v", err)
	// }
	//
	var comments []models.Comment
	// comments on posts
	for i, p := range posts {
		var c models.Comment

		c.Post_ID = p.ID
		c.Author_ID = users[i%len(users)].ID
		// pgtype.NullAssignTo(c.Parent_ID)
		c.Message = fmt.Sprintf("What a great post %s", users[i%len(users)].First_name)

		comments = append(comments, c)
	}

	batch := &pgx.Batch{}

	for _, c := range comments {
		batch.Queue(
			"INSERT INTO comments (post_id, author_id, message) VALUES ($1, $2, $3)",
			c.Post_ID,
			c.Author_ID,
			c.Message,
		)
	}

	br := pool.SendBatch(context.TODO(), batch)
	_, err = br.Exec()
	if err != nil {
		log.Fatalf("Failed to add comments to database: %v", err)
	}
}

func populateCommentsWithComments(pool *pgxpool.Pool) {
	users, err := database.GetAllUsers(pool)
	if err != nil {
		log.Fatalf("Failed to get users from database: %v", err)
	}
	comments, err := database.GetCommentsHelper(pool)
	if err != nil {
		log.Fatalf("Failed to get comments from database: %v", err)
	}

	for i, pc := range comments {
		var c models.Comment
		var pgID pgtype.Int8
		pgID.Int64Value()

		c.Post_ID = pc.Post_ID
		c.Author_ID = users[(i+1)%len(users)].ID
		c.Parent_ID = pgID
		c.Message = fmt.Sprintf("What a great comment %s", users[(i+1)%len(users)].First_name)
		comments = append(comments, c)
	}

	batch := &pgx.Batch{}

	for _, c := range comments {
		batch.Queue(
			"INSERT INTO comments (post_id, author_id, parent_id, message) VALUES ($1, $2, $3, $4)",
			c.Post_ID,
			c.Author_ID,
			c.Parent_ID,
			c.Message,
		)
	}

	br := pool.SendBatch(context.TODO(), batch)
	_, err = br.Exec()
	if err != nil {
		log.Fatalf("Failed to add comments to database: %v", err)
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

// func Int8(s int64) pgtype.Int8 {
// 	return pgtype.Int8{Int: s, Status: pgtype.Present}
// }
