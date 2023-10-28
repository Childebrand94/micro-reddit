package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/Childebrand94/micro-reddit/pkg/utils"
)

func AddUser(ctx context.Context, pool *pgxpool.Pool, user models.User) error {
	_, err := pool.Exec(
		ctx,
		"INSERT INTO users (first_name, last_name, username, email, password) VALUES ($1, $2, $3, $4, $5)",
		user.First_name,
		user.Last_name,
		user.Username,
		user.Email,
		user.Password,
	)

	return err
}

func GetAllUsers(ctx context.Context, pool *pgxpool.Pool) ([]models.User, error) {
	query := "SELECT * FROM users"
	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.First_name, &u.Last_name, &u.Username, &u.Email, &u.DateJoined); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func GetUserByID(ctx context.Context, pool *pgxpool.Pool, id int) ([]models.User, error) {
	userQuery := "SELECT * FROM users WHERE id = $1"

	var users []models.User
	var resp models.User

	// Fetch User
	row := pool.QueryRow(ctx, userQuery, id)
	err := row.Scan(
		&resp.ID,
		&resp.First_name,
		&resp.Last_name,
		&resp.Username,
		&resp.Email,
		&resp.DateJoined,
	)
	if err != nil {
		return nil, err
	}
	users = append(users, resp)

	return users, nil
}

func GetUserWithCPByID(ctx context.Context, pool *pgxpool.Pool, id int) (*models.UserResp, error) {
	userQuery := "SELECT * FROM users WHERE id = $1"
	postQuery := "SELECT * FROM posts WHERE author_id = $1"
	commentQuery := "SELECT * FROM comments WHERE author_id = $1"

	var resp models.UserResp

	// Fetch User
	row := pool.QueryRow(ctx, userQuery, id)
	err := row.Scan(
		&resp.User.ID,
		&resp.User.First_name,
		&resp.User.Last_name,
		&resp.User.Username,
		&resp.User.Email,
		&resp.User.DateJoined,
	)
	if err != nil {
		return nil, err
	}

	// Fetch Posts
	rows, err := pool.Query(ctx, postQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allPosts []models.Post

	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.ID,
			&post.Author_ID,
			&post.Title,
			&post.URL,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		totalVotes, err := utils.GetVoteTotal(pool, post.ID, "post_votes", "post_id")
		if err != nil {
			return nil, err
		}
		post.Vote = totalVotes

		allPosts = append(allPosts, post)
	}

	resp.Posts = utils.AddAuthorPosts(allPosts, resp.User)

	// Fetch Comments
	rows, err = pool.Query(ctx, commentQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.Post_ID,
			&comment.Author_ID,
			&comment.Parent_ID,
			&comment.Message,
			&comment.Created_at,
		)
		if err != nil {
			return nil, err
		}

		totalVotes, err := utils.GetVoteTotal(pool, comment.ID, "comment_vote", "comment_id")
		if err != nil {
			return nil, err
		}
		comment.Vote = totalVotes

		resp.Comments = append(resp.Comments, comment)
	}

	return &resp, nil
}

func UpdateUserByID(
	ctx context.Context,
	pool *pgxpool.Pool,
	updateUser models.User,
	id int64,
) error {
	query := `Update users SET first_name=$1, last_name=$2, username=$3, email=$4 WHERE id=$5`
	_, err := pool.Exec(
		ctx,
		query,
		updateUser.First_name,
		updateUser.Last_name,
		updateUser.Username,
		updateUser.Email,
		id,
	)
	return err
}
