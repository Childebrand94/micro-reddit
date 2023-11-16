package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// Load in .env file
	godotenv.Load(".env")

	// Set up database connection
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	// Ping the database by executing a trivial query
	err = dbpool.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping the database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully connected to the database!")

	// Read sql file
	sqlContent, err := os.ReadFile("./cmd/buildDb/micro-reddit-db-creation.sql")
	if err != nil {
		log.Fatalf("Error reading .sql file %v", err)
	}

	// Execute the SQL statement
	_, err = dbpool.Exec(context.Background(), string(sqlContent))
	if err != nil {
		log.Fatalf("Error executing SQL: %v", err)
	}

	fmt.Println("SQL executed successfully!")
}
