package database

import (
	"context"

	m "github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AddPostByUser(pool *pgxpool.Pool, author_id int64, title string, url string) error {
	_, err := pool.Exec(context.Background(), `INSERT INTO posts (author_id, url, title ) VALUES ($1, $2, $3)`, author_id, url, title)
	return err
}

func GetPostById(pool *pgxpool.Pool, post_id int64) (*m.Post, error) {
	var post m.Post
	err := pool.QueryRow(context.Background(), "SELECT author_id, title, url FROM posts").Scan(&post.Author_ID, &post.URL, &post.Title)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
