package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	"github.com/Childebrand94/micro-reddit/pkg/database"
	"github.com/Childebrand94/micro-reddit/pkg/models"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Load in .env file
	godotenv.Load(".env")

	// Set up database connection
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	populateDatabaseWithUsers(dbpool)
	println("Added Users")
	populateDatabaseWithPosts(dbpool)
	println("Added Posts")
	addVotesPosts(dbpool)
	println("Added Posts Votes")
	populateDatabaseWithComments(dbpool)
	println("Added Comments")
	populateCommentsWithComments(dbpool)
	println("Added Child Comments")
	populateCommentsWithVotes(dbpool)
	println("Added Comment votes")
	fmt.Println("Successfully populated database.")
}

func populateDatabaseWithUsers(pool *pgxpool.Pool) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte("1234"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	var users []models.User

	for i := 0; i < 30; i++ {
		var u models.User

		u.First_name = faker.FirstName()
		u.Last_name = faker.LastName()
		u.Username = faker.Username()
		u.Email = faker.Email()
		u.Password = string(hashedPassword)

		users = append(users, u)
	}
	batch := &pgx.Batch{}

	for _, u := range users {
		batch.Queue(
			"INSERT INTO users (first_name, last_name, username, email, password) VALUES ($1, $2, $3, $4, $5)",
			u.First_name,
			u.Last_name,
			u.Username,
			u.Email,
			u.Password,
		)
	}

	br := pool.SendBatch(context.TODO(), batch)
	_, err = br.Exec()
	if err != nil {
		log.Fatalf("Failed to add users to database: %v", err)
	}
}

func populateDatabaseWithPosts(pool *pgxpool.Pool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	users, err := database.GetAllUsers(ctx, pool)
	if err != nil {
		log.Fatalf("Failed to get users from database: %v", err)
	}
	var posts []models.Post

	for _, u := range users {
		var p models.Post

		p.Author_ID = u.ID
		p.URL = faker.URL()
		p.Title = faker.Sentence()

		posts = append(posts, p)
	}
	batch := &pgx.Batch{}

	for _, p := range posts {
		created_at := time.Now().Add(-time.Duration(rand.Intn(720)) * time.Hour)
		batch.Queue(
			"INSERT INTO posts (author_id, title, url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
			p.Author_ID,
			p.Title,
			p.URL,
			created_at,
			created_at,
		)
	}

	br := pool.SendBatch(ctx, batch)
	_, err = br.Exec()
	if err != nil {
		log.Fatalf("Failed to add posts to database: %v", err)
	}
}

func addVotesPosts(pool *pgxpool.Pool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	posts, err := database.GetPostsHelper(ctx, pool, "hot", "", nil)
	if err != nil {
		log.Fatalf("Failed to get posts from database: %v", err)
	}
	users, err := database.GetAllUsers(ctx, pool)
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

	batch := &pgx.Batch{}

	for _, v := range votes {
		created_at := time.Now().Add(-time.Duration(rand.Intn(720)) * time.Hour)
		batch.Queue(
			"INSERT INTO post_votes (user_id, post_id, up_vote, created_at) VALUES ($1, $2, $3, $4)",
			v.User_id,
			v.Post_id,
			v.Up_vote,
			created_at,
		)
	}

	br := pool.SendBatch(ctx, batch)
	_, err = br.Exec()
	if err != nil {
		log.Fatalf("Failed to add post votes to database: %v", err)
	}
}

func populateDatabaseWithComments(pool *pgxpool.Pool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	posts, err := database.GetPostsHelper(ctx, pool, "hot", "", nil)
	if err != nil {
		log.Fatalf("Failed to get posts from database: %v", err)
	}
	users, err := database.GetAllUsers(ctx, pool)
	if err != nil {
		log.Fatalf("Failed to get users from database: %v", err)
	}

	var comments []models.Comment
	commentsMessge := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed velit dignissim sodales ut eu sem integer vitae. Sed augue lacus viverra vitae congue eu consequat ac felis. Sed ullamcorper morbi tincidunt ornare massa eget. Viverra nam libero justo laoreet sit amet cursus. Sit amet consectetur adipiscing elit duis tristique. Interdum velit laoreet id donec ultrices tincidunt arcu non. Quis vel eros donec ac odio tempor orci dapibus ultrices. Nullam vehicula ipsum a arcu cursus vitae congue mauris rhoncus. Urna et pharetra pharetra massa massa ultricies mi quis hendrerit.`

	for i, p := range posts {
		var c models.Comment

		c.Post_ID = p.ID
		c.Author_ID = users[i%len(users)].ID
		c.Message = commentsMessge

		comments = append(comments, c)
	}

	tx, err := pool.Begin(ctx)
	if err != nil {
		log.Fatalf("Failed to start transaction: %v", err)
	}

	query := "INSERT INTO comments (post_id, author_id, message) VALUES ($1, $2, $3) RETURNING id"

	for _, c := range comments {
		var id int64

		err := tx.QueryRow(ctx, query, c.Post_ID, c.Author_ID, c.Message).Scan(&id)
		if err != nil {
			tx.Rollback(ctx)
			log.Fatalf("Failed to insert comment and retrieve ID: %v", err)
		}

		_, err = tx.Exec(ctx, "UPDATE comments SET path = CAST($1 as TEXT) WHERE id = $1 ", id)
		if err != nil {
			tx.Rollback(ctx)
			log.Fatalf("Failed to update path for root level comment: %v", err)
		}

	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}
}

func populateCommentsWithComments(pool *pgxpool.Pool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	users, err := database.GetAllUsers(ctx, pool)
	if err != nil {
		log.Fatalf("Failed to get users from database: %v", err)
	}
	parentComments, err := database.GetCommentsHelper(ctx, pool)
	if err != nil {
		log.Fatalf("Failed to get comments from database: %v", err)
	}

	var childComments []models.CommentResp
	message := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sed velit dignissim sodales ut eu sem integer vitae.`

	for i, pc := range parentComments {
		var c models.CommentResp

		c.Post_ID = pc.Post_ID
		c.Author_ID = users[(i+1)%len(users)].ID
		c.Parent_ID = sql.NullInt64{
			Int64: pc.ID,
			Valid: true,
		}
		c.Message = message
		c.Path = pc.Path

		childComments = append(childComments, c)
	}

	tx, err := pool.Begin(ctx)
	if err != nil {
		log.Fatalf("Failed to start transaction: %v", err)
	}

	query := "INSERT INTO comments (post_id, author_id, message, parent_id) VALUES ($1, $2, $3, $4) RETURNING id"

	for _, c := range childComments {
		var id int64

		err := tx.QueryRow(ctx, query, c.Post_ID, c.Author_ID, c.Message, c.Parent_ID).Scan(&id)
		if err != nil {
			tx.Rollback(ctx)
			log.Fatalf("Failed to insert child comment and retrieve ID: %v", err)
		}
		updateQuery := `UPDATE comments 
                        SET path = parent.path || '/' || CAST($2 AS TEXT) 
                        FROM comments AS parent 
                        WHERE parent.id = $1 AND comments.id = $2`

		_, err = tx.Exec(ctx, updateQuery, c.Parent_ID, id)
		if err != nil {
			tx.Rollback(ctx)
			log.Fatalf("Failed to update path for child comment: %v", err)
		}

	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func populateCommentsWithVotes(pool *pgxpool.Pool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	comments, err := database.GetCommentsHelper(ctx, pool)
	if err != nil {
		log.Fatalf("Failed to get users from database: %v", err)
	}

	users, err := database.GetAllUsers(ctx, pool)
	if err != nil {
		log.Fatalf("Failed to get users from database: %v", err)
	}

	type vote struct {
		User_id    int64
		Comment_id int64
		Up_vote    bool
	}

	var votes []vote

	for _, p := range comments {
		for _, u := range users {
			var v vote

			v.User_id = u.ID
			v.Comment_id = p.ID
			v.Up_vote = randomBool()

			votes = append(votes, v)
		}
	}
	batch := &pgx.Batch{}

	for _, v := range votes {
		batch.Queue(
			"INSERT INTO comment_votes (user_id, comment_id, up_vote) VALUES ($1, $2, $3)",
			v.User_id,
			v.Comment_id,
			v.Up_vote,
		)
	}

	br := pool.SendBatch(ctx, batch)
	err = br.Close()
	if err != nil {
		log.Fatalf("Failed to add comments to database: %v", err)
	}
}
