package main

import (
	"net/http"
	"time"

	"github.com/orange432/bloggo/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// routes handles all of the routes in our todo list
func routes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	// Timeout
	router.Use(middleware.Timeout(120 * time.Second))

	// Standard page routes
	router.Get("/", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)
	router.Get("/test", handlers.Repo.TestPage)
	router.Get("/login", handlers.Repo.Login)
	router.Post("/login", handlers.Repo.LoginPost)
	router.Route("/articles", func(router chi.Router) {
		router.Get("/{articleSlug}", handlers.Repo.GetArticle)
	})

	fileServer := http.FileServer(http.Dir("./public/"))

	router.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return router
}
