package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
  router http.Handler
  DB *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *App {
  app := &App{
		  DB: pool,
	}
  app.loadRoutes()

  return app
}

func (a *App) Start(ctx context.Context) error {
  server := &http.Server{
    Addr : ":3000",
        Handler: a.router,
  }

  

  err := server.ListenAndServe()
  if err != nil {
    return fmt.Errorf("failed to start server: %w", err)
  }
  return nil 
}
