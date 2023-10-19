package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/Childebrand94/micro-reddit/pkg/utils"
)

func AddPostByUser(pool *pgxpool.Pool, author_id int64, url string) error {
	_, err := pool.Exec(
		context.TODO(),
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

	totalVotes, err := utils.GetVoteTotal(pool, p.ID, "post_votes", "post_id")
	if err != nil {
		return nil, nil, err
	}
	p.Vote = totalVotes
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
		totalVotes, err := utils.GetVoteTotal(pool, c.ID, "comment_vote", "comment_id")
		if err != nil {
			return nil, nil, err
		}
		c.Vote = totalVotes

		comments = append(comments, c)
	}
	defer rows.Close()

	return posts, comments, nil
}

func GetAllPosts(pool *pgxpool.Pool) ([]models.Comment, []models.Post, error) {
	allPosts, err := GetPostsHelper(pool)
	if err != nil {
		return nil, nil, err
	}

	allComments, err := GetCommentsHelper(pool)
	if err != nil {
		return nil, nil, err
	}

	return allComments, allPosts, nil
}

func AddPostVotes(pool *pgxpool.Pool, user_id, post_id int64, up_vote bool) error {
	query := "INSERT INTO up_vote (user_id, post_id, up_vote) VALUES ($1, $2, $3)"
	_, err := pool.Exec(context.TODO(), query, user_id, post_id, up_vote)
	return err
}

func GetPostsHelper(pool *pgxpool.Pool) ([]models.Post, error) {
	queryForPosts := `SELECT * FROM posts;`

	postRows, err := pool.Query(context.TODO(), queryForPosts)
	if err != nil {
		return nil, err
	}

	defer postRows.Close()

	var allPosts []models.Post

	for postRows.Next() {
		// Create structs to scan data too
		var p models.Post

		// Scan data from DB to structs
		if err := postRows.Scan(
			&p.ID,
			&p.Author_ID,
			&p.URL,
			&p.CreatedAt,
			&p.UpdatedAt); err != nil {
			return nil, err
		}
		p.Vote, err = utils.GetVoteTotal(pool, p.ID, "post_votes", "post_id")
		if err != nil {
			return nil, err
		}
		allPosts = append(allPosts, p)
	}

	return allPosts, nil
}

func GetCommentsHelper(pool *pgxpool.Pool) ([]models.Comment, error) {
	queryForComments := `SELECT * FROM comments;`

	commentRows, err := pool.Query(context.TODO(), queryForComments)
	if err != nil {
		return nil, err
	}
	defer commentRows.Close()

	var allComments []models.Comment

	for commentRows.Next() {
		var c models.Comment

		if err := commentRows.Scan(
			&c.ID,
			&c.Post_ID,
			&c.Author_ID,
			&c.Parent_ID,
			&c.Message,
			&c.Created_at,
		); err != nil {
			return nil, err
		}
		totalVotes, err := utils.GetVoteTotal(pool, c.ID, "comment_vote", "comment_id")
		if err != nil {
			return nil, err
		}
		c.Vote = totalVotes

		allComments = append(allComments, c)
	}

	return allComments, nil
}
