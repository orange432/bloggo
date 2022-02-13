package main

import (
	"bloggo/pkg/config"
	"bloggo/pkg/handlers"
	"bloggo/pkg/render"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var PORT = ":8000"

func main() {
	var app config.AppConfig

	tCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Couldn't load template cache")
	}
	app.TemplateCache = tCache
	render.NewTemplates(&app)

	r := mux.NewRouter()

	// Page Routes
	r.HandleFunc("/", handlers.Home).Methods("GET")
	r.HandleFunc("/login", handlers.LoginPage).Methods("GET")
	r.HandleFunc("/register", handlers.RegisterPage).Methods("GET")
	r.HandleFunc("/dashboard", handlers.Dashboard).Methods("GET")
	r.HandleFunc("/editor", handlers.EditorPage).Methods("GET")
	r.HandleFunc("/articles", handlers.Home).Methods("GET")
	r.HandleFunc("/articles/{id}", handlers.ArticlePage).Methods("GET")

	// API routes
	r.HandleFunc("/api/login", handlers.Login).Methods("POST")
	r.HandleFunc("/api/register", handlers.Register).Methods("POST")
	r.HandleFunc("/api/editor", handlers.Editor).Methods("POST")
	r.HandleFunc("/api/articles", handlers.ListArticles).Methods("GET")

	fileServer := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/").Handler(http.StripPrefix("/public", fileServer))

	fmt.Println(fmt.Sprintf("🚀 Running at http://localhost%s", PORT))
	log.Fatal(http.ListenAndServe(PORT, r))
}
