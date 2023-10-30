package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/Childebrand94/micro-reddit/pkg/utils"
)

func AddPostByUser(
	ctx context.Context,
	pool *pgxpool.Pool,
	author_id int64,
	url, title string,
) error {
	_, err := pool.Exec(
		ctx,
		`INSERT INTO posts (author_id, url, title) VALUES ($1, $2, $3)`,
		author_id,
		title,
		url,
	)
	return err
}

func GetPostById(
	ctx context.Context,
	pool *pgxpool.Pool,
	post_id int64,
) ([]models.Post, []models.CommentResp, error) {
	var posts []models.Post
	var p models.Post
	var comments []models.CommentResp

	queryPosts := "SELECT * FROM posts WHERE id = $1"
	queryComments := `select c.id, c.post_id, c.author_id, c.parent_id, c.message, c.created_at, u.first_name, u.last_name, u.username 
						from "comments" c  
						left join 
						users u on u.id  = c.author_id 
						where post_id = $1`

	row := pool.QueryRow(ctx, queryPosts, post_id)
	if err := row.Scan(&p.ID, &p.Author_ID, &p.Title, &p.URL, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return nil, nil, err
	}

	totalVotes, err := utils.GetVoteTotal(pool, p.ID, "post_votes", "post_id")
	if err != nil {
		return nil, nil, err
	}
	p.Vote = totalVotes
	posts = append(posts, p)

	rows, err := pool.Query(ctx, queryComments, post_id)
	if err != nil {
		return nil, nil, err
	}

	for rows.Next() {
		var c models.CommentResp
		if err := rows.Scan(&c.ID, &c.Post_ID, &c.Author_ID, &c.Parent_ID, &c.Message, &c.Created_at, &c.Author.FirstName, &c.Author.LastName, &c.Author.UserName); err != nil {
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

func GetAllPosts(ctx context.Context, pool *pgxpool.Pool) ([]models.PostResponse, error) {
	allPosts, err := GetPostsHelper(ctx, pool)
	if err != nil {
		return nil, err
	}

	var postResp []models.PostResponse

	for _, post := range allPosts {
		query := "SELECT * FROM comments WHERE post_id = $1"

		rows, err := pool.Query(ctx, query, post.ID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var c models.CommentResp
			if err := rows.Scan(&c.ID, &c.Post_ID, &c.Author_ID, &c.Parent_ID, &c.Message, &c.Created_at); err != nil {
				return nil, err
			}
			post.Comments = append(post.Comments, c)
		}
		postResp = append(postResp, post)
	}

	return postResp, nil
}

func AddPostVotes(
	ctx context.Context,
	pool *pgxpool.Pool,
	user_id, post_id int64,
	up_vote bool,
) error {
	query := "INSERT INTO post_votes (user_id, post_id, up_vote) VALUES ($1, $2, $3)"
	_, err := pool.Exec(ctx, query, user_id, post_id, up_vote)
	return err
}

func GetPostsHelper(ctx context.Context, pool *pgxpool.Pool) ([]models.PostResponse, error) {
	queryForPosts := `SELECT p.id, p.author_id, p.title, p.url, p.created_at, p.updated_at, u.first_name,
                        u.last_name, u.username
                    FROM posts as p
                    LEFT JOIN 
                    users as u ON u.id = p.author_id`

	postRows, err := pool.Query(ctx, queryForPosts)
	if err != nil {
		return nil, err
	}

	defer postRows.Close()

	var allPosts []models.PostResponse

	for postRows.Next() {

		var p models.PostResponse

		if err := postRows.Scan(
			&p.ID,
			&p.Author_ID,
			&p.Title,
			&p.URL,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.Author.FirstName,
			&p.Author.LastName,
			&p.Author.UserName); err != nil {
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

func GetCommentsHelper(ctx context.Context, pool *pgxpool.Pool) ([]models.CommentResp, error) {
	queryForComments := `select c.id, c.post_id, c.author_id, c.parent_id, c.message, c.created_at, u.first_name, u.last_name, u.username 
						from "comments" c  
						left join 
						users u on u.id  = c.author_id `

	commentRows, err := pool.Query(ctx, queryForComments)
	if err != nil {
		return nil, err
	}
	defer commentRows.Close()

	var allComments []models.CommentResp

	for commentRows.Next() {
		var c models.CommentResp

		if err := commentRows.Scan(
			&c.ID,
			&c.Post_ID,
			&c.Author_ID,
			&c.Parent_ID,
			&c.Message,
			&c.Created_at,
			&c.Author.FirstName,
			&c.Author.LastName,
			&c.Author.UserName,
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
