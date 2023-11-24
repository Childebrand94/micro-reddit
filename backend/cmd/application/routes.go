package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/Childebrand94/micro-reddit/pkg/handler"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	// Setup middleware
	cors := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(middleware.Logger)
	router.Use(cors.Handler)

	router.Route("/posts", a.loadPostRoutes)
	router.Route("/users", a.loadUserRoutes)
	router.Route("/sessions", a.loadSessionRoutes)
	a.router = router
}

func (a *App) loadPostRoutes(router chi.Router) {
	postHandler := &handler.Post{
		DB: a.DB,
	}
	commentHandler := &handler.Comment{
		DB: a.DB,
	}

	router.Post("/", postHandler.Create)
	router.Get("/", postHandler.List)
	router.Get("/{id}", postHandler.GetByID)
	router.Put("/{id}/{vote}", postHandler.PostVotes)
	router.Put("/{id}", postHandler.UpdateByID)
	router.Delete("/{id}", postHandler.DeleteByID)
	router.Post("/{id}/comments", commentHandler.Create)
	router.Get("/{id}/comments", commentHandler.List)
	router.Put("/{post_id}/comments/{comment_id}/{vote}", commentHandler.CommentVotes)
}

func (a *App) loadUserRoutes(router chi.Router) {
	userHandler := &handler.User{
		DB: a.DB,
	}

	router.Get("/", userHandler.List)
	router.Post("/", userHandler.Create)
	router.Put("/", userHandler.UpdateByID)
	router.Post("/login", userHandler.Authenticate)
	router.Post("/logout", userHandler.Logout)
	router.Get("/{id}", userHandler.GetByID)
	router.Get("/{id}/posts", userHandler.GetAllPostsByUser)
	router.Get("/{id}/comments", userHandler.GetAllCommentsByUser)
	router.Get("/{id}/points", userHandler.GetUserPoints)
}

func (a *App) loadSessionRoutes(router chi.Router) {
	sessionHandler := &handler.Session{
		DB: a.DB,
	}

	router.Get("/", sessionHandler.IsSession)
}
