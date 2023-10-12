package database

import (
	"context"

	m "github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AddUser(pool *pgxpool.Pool, fName string, lName string, email string) error {
	_, err := pool.Exec(context.Background(), "INSERT INTO users (first_name, last_name, email) VALUES ($1, $2, $3)", fName, lName, email)
	return err
}

func GetAllUsers(pool *pgxpool.Pool) ([]m.User, error) {
	query := "SELECT * FROM users"
    rows, err := pool.Query(context.Background(), query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []m.User
    for rows.Next() {
        var u m.User
        if err := rows.Scan(&u.ID, &u.First_name, &u.Last_name,&u.Email,&u.DateJoined); err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    return users, nil
}
