package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/Childebrand94/micro-reddit/pkg/utils"
)

func AddUser(ctx context.Context, pool *pgxpool.Pool, user models.User) (int64, error) {
	var userID int
	err := pool.QueryRow(
		ctx,
		"INSERT INTO users (first_name, last_name, username, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		user.First_name,
		user.Last_name,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return int64(userID), nil
}

func GetAllUsers(ctx context.Context, pool *pgxpool.Pool) ([]models.UserResp, error) {
	query := "SELECT id, first_name, last_name, username, email, registered_at FROM users"
	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserResp
	for rows.Next() {
		var u models.UserResp
		if err := rows.Scan(&u.ID, &u.First_name, &u.Last_name, &u.Username, &u.Email, &u.DateJoined); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func GetUserByID(ctx context.Context, pool *pgxpool.Pool, id int) (*models.UserResp, error) {
	userQuery := "SELECT id, first_name, last_name, username, email, registered_at FROM users WHERE id = $1"

	var resp models.UserResp

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

	return &resp, nil
}

func UpdateUserPassword(
	ctx context.Context,
	pool *pgxpool.Pool,
	newPassword string,
	id int64,
) error {
	query := `UPDATE users SET password = $1 WHERE id = $2`
	_, err := pool.Exec(
		ctx,
		query,
		newPassword,
		id,
	)
	return err
}

func GetUserByEmail(ctx context.Context, pool *pgxpool.Pool, email string) (*models.User, error) {
	query := `SELECT * FROM users WHERE email = $1`

	var u models.User
	row := pool.QueryRow(ctx, query, email)
	err := row.Scan(
		&u.ID,
		&u.First_name,
		&u.Last_name,
		&u.Username,
		&u.Email,
		&u.Password,
		&u.DateJoined,
	)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func CreateSession(ctx context.Context, pool *pgxpool.Pool, sessionId string, userId int64) error {
	query := "INSERT INTO sessions (session_id, user_id) VALUES ($1, $2)"

	_, err := pool.Exec(ctx, query, sessionId, userId)
	return err
}

func DeleteSession(ctx context.Context, pool *pgxpool.Pool, sessionId string) error {
	query := `DELETE FROM sessions WHERE session_id = $1`
	_, err := pool.Exec(ctx, query, sessionId)
	return err
}

func GetAllPostsByUser(ctx context.Context, pool *pgxpool.Pool, id int64, votersId *int64) ([]models.PostResponse, error) {
	query := `
		SELECT 
			p.id, p.author_id, p.title, p.url, p.created_at, p.updated_at,
			u.first_name, u.last_name, u.username 
		FROM posts p
		LEFT JOIN users u ON p.author_id = u.id
		WHERE p.author_id = $1
	`

	rows, err := pool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.PostResponse

	for rows.Next() {
		var p models.PostResponse
		if err := rows.Scan(
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

		p.UsersVoteStatus, err = utils.UserPostVoteCheck(ctx, pool, p.ID, votersId)
		if err != nil {
			return nil, err
		}

		queryComments := `SELECT * 
                        FROM comments c
                        WHERE c.post_id = $1`

		rows, err := pool.Query(ctx, queryComments, p.ID)
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
				&c.Path,
				&c.Created_at,
			); err != nil {
				return nil, err
			}
			p.Comments = append(p.Comments, c)
		}

		posts = append(posts, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func GetAllCommentsByUser(ctx context.Context, pool *pgxpool.Pool, id int64) ([]models.CommentResp, error) {
	query := `
    SELECT 
        c.id, 
        c.post_id, 
        c.author_id, 
        c.parent_id, 
        c.message, 
        c.created_at, 
        u.first_name, 
        u.last_name, 
        u.username 
    FROM "comments" AS c  
    LEFT JOIN users AS u ON u.id = c.author_id
    WHERE c.author_id = $1`

	rows, err := pool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.CommentResp

	for rows.Next() {
		var c models.CommentResp

		if err := rows.Scan(&c.ID, &c.Post_ID, &c.Author_ID, &c.Parent_ID, &c.Message,
			&c.Created_at, &c.Author.FirstName, &c.Author.LastName, &c.Author.UserName); err != nil {
			return nil, err
		}

		totalVotes, err := utils.GetVoteTotal(pool, c.ID, "comment_votes", "comment_id")
		if err != nil {
			return nil, err
		}
		c.Vote = totalVotes

		comments = append(comments, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func GetUserPoints(ctx context.Context, pool *pgxpool.Pool, id int64) (*models.UserPoints, error) {
	queryPosts := `
        SELECT 
            COALESCE(COUNT(p.id), 0) AS post_count
        FROM 
            posts p
        WHERE 
            p.author_id = $1;`

	queryPostVotes := `
        SELECT 
            COALESCE(SUM(CASE WHEN pv.up_vote = true THEN 1 ELSE 0 END), 0) AS up_vote_count,
            COALESCE(SUM(CASE WHEN pv.up_vote = false THEN 1 ELSE 0 END), 0) AS down_vote_count
        FROM 
            post_votes AS pv 
        WHERE 
            pv.user_id = $1;`

	var up models.UserPoints

	err := pool.QueryRow(ctx, queryPosts, id).Scan(&up.PostCount)
	if err != nil {
		return nil, err
	}

	err = pool.QueryRow(ctx, queryPostVotes, id).Scan(&up.PostUpVotes, &up.PostDownVotes)
	if err != nil {
		return nil, err
	}

	return &up, nil
}

func EmailExists(ctx context.Context, pool *pgxpool.Pool, email string) (bool, error) {
	var exists bool
	err := pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
