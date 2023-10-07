package main

// import (
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/chi/v5/middleware"
// )

// func main() {
// 	r := chi.NewRouter()

// 	dbpool, err := ConnDb()
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Failed to conned to DB: %n \n", err)
// 		os.Exit(1)
// 	}
// 	defer dbpool.Close()

// 	r.Use(middleware.Logger)
// 	r.Use(middleware.Recoverer)

// 	http.ListenAndServe(":8080", r)
// }
