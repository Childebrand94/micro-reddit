package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Childebrand94/micro-reddit/pkg/models"
)

func AddUser(ctx context.Context, pool *pgxpool.Pool, user models.User) error {
	_, err := pool.Exec(
		ctx,
		"INSERT INTO users (first_name, last_name, username, email) VALUES ($2, $2, $3, $4)",
		user.First_name,
		user.Last_name,
		user.Username,
		user.Email,
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
	query := "SELECT * FROM users WHERE id = $1"
	row := pool.QueryRow(ctx, query, id)
	var users []models.User

	var user models.User
	err := row.Scan(
		&user.ID,
		&user.First_name,
		&user.Last_name,
		&user.Username,
		&user.Email,
		&user.DateJoined,
	)
	if err != nil {
		return nil, err
	}
	users = append(users, user)

	return users, nil
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
