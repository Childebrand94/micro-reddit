package database

import (
	"context"

	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AddPostByUser(pool *pgxpool.Pool, author_id int64, url string) error {
	_, err := pool.Exec(context.Background(), `INSERT INTO posts (author_id, url) VALUES ($1, $2)`, author_id, url)
	return err
}

func GetPostById(pool *pgxpool.Pool, post_id int64) (*models.Post, error) {
	var post models.Post
	err := pool.QueryRow(context.Background(), "SELECT author_id, url FROM posts").Scan(&post.Author_ID, &post.URL)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func GetAllPosts(pool *pgxpool.Pool) (map[int64]models.PostWithComments, error) {
	query :=
		`SELECT 
		posts.id, 
		posts.author_id, 
		posts.url, 
		posts.created_at, 
		posts.updated_at, 
		comments.id, 
		comments.post_id, 
		comments.author_id, 
		comments.parent_id, 
		comments.message, 
		comments.created_at 
	FROM posts 
	LEFT JOIN comments ON posts.id = comments.post_id;`

	rows, err := pool.Query(context.TODO(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	combinedMap := make(map[int64]models.PostWithComments)

	for rows.Next() {
		//Create structs to scan data too
		var comment models.Comment
		var post models.Post

		// Scan data from DB to structs
		if err := rows.Scan(
			&post.ID,
			&post.Author_ID,
			&post.URL,
			&post.CreatedAt,
			&post.UpdatedAt,
			&comment.ID,
			&comment.Post_ID,
			&comment.Author_ID,
			&comment.Parent_ID,
			&comment.Message,
			&comment.Created_at); err != nil {
			return nil, err
		}

		//Check if post is already present in postDataWithComments
		//If present append the comments to Comments slice

		if postDataWithComments, exists := combinedMap[post.ID]; exists {
			if comment.ID.Valid {
				println(comment.ID.Valid)
				postDataWithComments.Comments = append(postDataWithComments.Comments, comment)
				combinedMap[post.ID] = postDataWithComments
			}
			// If no post was present create a new key in combined map and reference this post
		} else {
			combinedMap[post.ID] = models.PostWithComments{
				PostData: post,
				Comments: func() []models.Comment {
					if comment.ID.Valid {
						return []models.Comment{comment}
					}
					return []models.Comment{}
				}(),
			}
		}
	}
	return combinedMap, nil
}
