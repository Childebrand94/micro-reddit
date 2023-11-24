package application

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
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
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	server := &http.Server{
		Addr:    ":" + port,
		Handler: a.router,
	}
	err := a.DB.Ping(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	defer func() {
		a.DB.Close()
		fmt.Println("Closed PostgreSQL connection", err)
	}()

	// Creating channel
	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	fmt.Println("Starting server...")

	// give the server a moment to start
	time.Sleep(time.Second)

	// check if the database is reachable
	err = a.DB.Ping(ctx)
	if err != nil {
		return fmt.Errorf("failed to ping postgresql: %w", err)
	}

	fmt.Println("Server is running and connected to PostgreSQL on port", port)

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
