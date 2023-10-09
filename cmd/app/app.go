package main

import (
  "os"
  "fmt"
  "context"

  "github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)


type App struct{
  Router *chi.Mux
  DB *pgxpool.Pool
}

func (a *App) Initalize() {
  dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
  	if err != nil {
  		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
  		os.Exit(1)
	}
}
