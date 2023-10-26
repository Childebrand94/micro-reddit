package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	router http.Handler
	DB     *pgxpool.Pool
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
		Addr:    ":3000",
		Handler: a.router,
	}
	err := a.DB.Ping(ctx)
	if err != nil {
		return fmt.Errorf("Failed to connect to PostgreSQL: %w", err)
	}

	defer func() {
		a.DB.Close()
		fmt.Println("Closed PostgreSQL connection", err)
	}()

	fmt.Println("Starting server...")

	// creating channel
	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("Failed to start server: %w", err)
		}
		close(ch)
	}()

	// setting up receiver for channel
	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
