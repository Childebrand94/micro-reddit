package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func AddUser(pool *pgxpool.Pool, fName string, lName string, email string) error {
	_, err := pool.Exec(context.Background(), "INSERT INTO users (first_name, last_name, email) VALUES ($1, $2, $3)", fName, lName, email)
	return err
}