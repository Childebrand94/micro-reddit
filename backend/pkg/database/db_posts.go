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
		`INSERT INTO posts (author_id, title, url) VALUES ($1, $2, $3)`,
		author_id,
		title,
		url,
	)
	return err
}

func GetPostById(ctx context.Context, pool *pgxpool.Pool, post_id int64, userId *int64) (*models.PostResponse, error) {
	queryPosts := `SELECT 
                        p.id, 
                        p.author_id, 
                        p.title, 
                        p.url, 
                        p.created_at, 
                        p.updated_at, 
                        u.first_name, 
                        u.last_name, 
                        u.username
                    FROM 
                        posts AS p
                        LEFT JOIN users AS u ON u.id = p.author_id
                    WHERE 
                        p.id = $1`

	queryComments := `SELECT 
                        c.id, 
                        c.post_id, 
                        c.author_id, 
                        c.parent_id, 
                        c.message, 
                        c.created_at, 
                        u.first_name, 
                        u.last_name, 
                        u.username
                    FROM 
                        comments AS c
                        LEFT JOIN users AS u ON u.id = c.author_id
                    WHERE 
                        post_id = $1`

	var p models.PostResponse
	row := pool.QueryRow(ctx, queryPosts, post_id)
	if err := row.Scan(
		&p.ID,
		&p.Author_ID,
		&p.Title,
		&p.URL,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Author.FirstName,
		&p.Author.LastName,
		&p.Author.UserName,
	); err != nil {
		return nil, err
	}

	totalVotes, err := utils.GetVoteTotal(pool, p.ID, "post_votes", "post_id")
	if err != nil {
		return nil, err
	}
	p.Vote = totalVotes

	p.UsersVoteStatus, err = utils.UserPostVoteCheck(ctx, pool, p.ID, userId)
	if err != nil {
		return nil, err
	}

	rows, err := pool.Query(ctx, queryComments, post_id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c models.CommentResp
		if err := rows.Scan(
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
		totalVotes, err := utils.GetVoteTotal(pool, c.ID, "comment_votes", "comment_id")
		if err != nil {
			return nil, err
		}
		c.Vote = totalVotes
		c.UsersVoteStatus, err = utils.UserCommentVoteCheck(ctx, pool, c.ID, userId)
		if err != nil {
			return nil, err
		}

		p.Comments = append(p.Comments, c)
	}
	defer rows.Close()

	return &p, nil
}

func GetAllPosts(ctx context.Context, pool *pgxpool.Pool, sort, search string, userId *int64) ([]models.PostResponse, error) {
	allPosts, err := GetPostsHelper(ctx, pool, sort, search, userId)
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
			if err := rows.Scan(
				&c.ID,
				&c.Post_ID,
				&c.Author_ID,
				&c.Parent_ID,
				&c.Message,
				&c.Created_at,
			); err != nil {
				return nil, err
			}
			c.UsersVoteStatus, err = utils.UserCommentVoteCheck(ctx, pool, c.ID, userId)
			if err != nil {
				return nil, err
			}

			post.Comments = append(post.Comments, c)
		}
		postResp = append(postResp, post)

	}

	return postResp, nil
}

func AddPostVotes(ctx context.Context, pool *pgxpool.Pool, user_id, post_id int64, up_vote bool) error {
	query := `INSERT INTO post_votes (post_id, user_id, up_vote)
                VALUES ($1, $2, $3)
                ON CONFLICT (post_id, user_id)
                DO UPDATE SET up_vote = EXCLUDED.up_vote;`
	_, err := pool.Exec(ctx, query, post_id, user_id, up_vote)
	return err
}

func GetPostsHelper(ctx context.Context, pool *pgxpool.Pool, sort, search string, userId *int64) ([]models.PostResponse, error) {
	if search == "" {
		allPosts, err := GetPostsSort(ctx, pool, sort, userId)
		if err != nil {
			return nil, err
		}
		return allPosts, nil
	} else {
		allPosts, err := GetPostsSearch(ctx, pool, search, userId)
		if err != nil {
			return nil, err
		}
		return allPosts, nil
	}
}

func GetPostsSort(ctx context.Context, pool *pgxpool.Pool, sort string, userId *int64) ([]models.PostResponse, error) {
	query := utils.GetSortMethod(sort)

	postRows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer postRows.Close()

	var allPosts []models.PostResponse

	for postRows.Next() {

		var p models.PostResponse
		var totalvotes int

		if err := postRows.Scan(
			&p.ID,
			&p.Author_ID,
			&p.Title,
			&p.URL,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.Author.FirstName,
			&p.Author.LastName,
			&p.Author.UserName,
			&totalvotes); err != nil {
			return nil, err
		}
		p.Vote, err = utils.GetVoteTotal(pool, p.ID, "post_votes", "post_id")
		if err != nil {
			return nil, err
		}
		p.UsersVoteStatus, err = utils.UserPostVoteCheck(ctx, pool, p.ID, userId)
		if err != nil {
			return nil, err
		}

		allPosts = append(allPosts, p)
	}

	return allPosts, nil
}

func GetPostsSearch(ctx context.Context, pool *pgxpool.Pool, search string, userId *int64) ([]models.PostResponse, error) {
	searchTerm := "%" + search + "%"
	query := `SELECT 
                p.id,
                p.author_id,
                p.title,
                p.url,
                p.created_at,
                p.updated_at,
                u.first_name,
                u.last_name,
                u.username
            FROM 
                posts p
            LEFT JOIN 
                users u ON u.id = p.author_id
            WHERE p.title ILIKE $1
                OR u.username ILIKE $1
                OR p.url ILIKE $1;`

	postRows, err := pool.Query(ctx, query, searchTerm)
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
		p.UsersVoteStatus, err = utils.UserPostVoteCheck(ctx, pool, p.ID, userId)
		if err != nil {
			return nil, err
		}

		allPosts = append(allPosts, p)
	}
	return allPosts, nil
}

func GetCommentsHelper(ctx context.Context, pool *pgxpool.Pool) ([]models.CommentResp, error) {
	queryForComments := `SELECT 
                            c.id, 
                            c.post_id, 
                            c.author_id, 
                            c.parent_id, 
                            c.message, 
                            c.path,
                            c.created_at, 
                            u.first_name, 
                            u.last_name, 
                            u.username 
                        FROM "comments" AS c  
                        LEFT JOIN users AS u ON u.id = c.author_id`

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
			&c.Path,
			&c.Created_at,
			&c.Author.FirstName,
			&c.Author.LastName,
			&c.Author.UserName,
		); err != nil {
			return nil, err
		}
		totalVotes, err := utils.GetVoteTotal(pool, c.ID, "comment_votes", "comment_id")
		if err != nil {
			return nil, err
		}
		c.Vote = totalVotes

		allComments = append(allComments, c)
	}

	return allComments, nil
}
