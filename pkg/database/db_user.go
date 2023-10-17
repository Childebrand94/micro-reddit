package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	m "github.com/Childebrand94/micro-reddit/pkg/models"
)

func AddUser(pool *pgxpool.Pool, fName, lName, username, email string) error {
	_, err := pool.Exec(
		context.TODO(),
		"INSERT INTO users (first_name, last_name, username, email) VALUES ($1, $2, $3, $4)",
		fName,
		lName,
		username,
		email,
	)

	return err
}

func GetAllUsers(pool *pgxpool.Pool) ([]m.User, error) {
	query := "SELECT * FROM users"
	rows, err := pool.Query(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []m.User
	for rows.Next() {
		var u m.User
		if err := rows.Scan(&u.ID, &u.First_name, &u.Last_name, &u.Username, &u.Email, &u.DateJoined); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func GetUserByID(pool *pgxpool.Pool, id int) (*m.User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	row := pool.QueryRow(context.TODO(), query, id)

	var user m.User
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
	return &user, nil
}

// func UpdateUserByID(pool *pgxpool.Pool, id int) (*m.User, error) {
// query := "UPDATE users SET name=$1 WHERE id=$2"
// row := pool.QueryRow(context.TODO(),query,id)
// }