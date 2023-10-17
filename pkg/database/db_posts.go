package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/models"
)

func AddPostByUser(pool *pgxpool.Pool, author_id int64, url string) error {
	_, err := pool.Exec(
		context.Background(),
		`INSERT INTO posts (author_id, url) VALUES ($1, $2)`,
		author_id,
		url,
	)
	return err
}

func GetPostById(pool *pgxpool.Pool, post_id int64) ([]models.Post, []models.Comment, error) {
	var posts []models.Post
	var p models.Post
	var comments []models.Comment

	queryPosts := "SELECT * FROM posts WHERE id = $1"
	queryComments := "SELECT * FROM comments WHERE post_id = $1"

	row := pool.QueryRow(context.TODO(), queryPosts, post_id)
	if err := row.Scan(&p.ID, &p.Author_ID, &p.URL, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return nil, nil, err
	}

	posts = append(posts, p)

	rows, err := pool.Query(context.TODO(), queryComments, post_id)
	if err != nil {
		return nil, nil, err
	}

	for rows.Next() {
		var c models.Comment
		if err := rows.Scan(&c.ID, &c.Post_ID, &c.Author_ID, &c.Parent_ID, &c.Message, &c.Created_at); err != nil {
			return nil, nil, err
		}
		comments = append(comments, c)
	}
	defer rows.Close()

	return posts, comments, nil
}

func GetAllPosts(pool *pgxpool.Pool) ([]models.Comment, []models.Post, error) {
	queryForPosts := `SELECT * FROM posts;`

	queryForComments := `SELECT * FROM comments;`

	postRows, err := pool.Query(context.TODO(), queryForPosts)
	if err != nil {
		return nil, nil, err
	}

	defer postRows.Close()

	var allPosts []models.Post

	for postRows.Next() {
		// Create structs to scan data too
		var post models.Post

		// Scan data from DB to structs
		if err := postRows.Scan(
			&post.ID,
			&post.Author_ID,
			&post.URL,
			&post.CreatedAt,
			&post.UpdatedAt); err != nil {
			return nil, nil, err
		}
		allPosts = append(allPosts, post)
	}

	commentRows, err := pool.Query(context.TODO(), queryForComments)
	if err != nil {
		return nil, nil, err
	}

	defer commentRows.Close()

	var allComments []models.Comment

	for commentRows.Next() {
		var comment models.Comment

		if err := commentRows.Scan(
			&comment.ID,
			&comment.Post_ID,
			&comment.Author_ID,
			&comment.Parent_ID,
			&comment.Message,
			&comment.Created_at,
		); err != nil {
			return nil, nil, err
		}
		allComments = append(allComments, comment)
	}
	return allComments, allPosts, nil
}