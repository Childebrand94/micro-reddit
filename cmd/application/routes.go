package application

import (
	"net/http"

	"github.com/Childebrand94/micro-reddit/pkg/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func (a *App)loadRoutes() {
  router := chi.NewRouter()

  router.Use(middleware.Logger)

  router.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome to Reddit"))
    w.WriteHeader(http.StatusOK)
  })

  router.Route("/posts", a.loadPostRoutes)
  router.Route("/users", a.loadUserRoutes)
  
  a.router = router
}

func (a *App) loadPostRoutes(router chi.Router) {
  postHandler := &handler.Post{
    DB: a.DB, 
  }

  router.Post("/", postHandler.Create)
  router.Get("/", postHandler.List)
  router.Get("/{id}", postHandler.GetByID)
  router.Put("/{id}", postHandler.UpdateByID)
  router.Delete("/{id}", postHandler.DeleteByID)
}

func (a *App) loadUserRoutes(router chi.Router) {
  userHandler := &handler.User{
    DB: a.DB,
  }

  router.Post("/",userHandler.Create)
  router.Get("/", userHandler.List)
  router.Get("/{id}", userHandler.GetByID)
  router.Put("/{id}", userHandler.UpdateByID)
  router.Delete("/{id}", userHandler.DeleteByID)
}
