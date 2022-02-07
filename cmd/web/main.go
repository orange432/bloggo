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

	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/api/save-article", handlers.SaveArticle).Methods("POST")

	fmt.Println(fmt.Sprintf("🚀 Running at http://localhost%s", PORT))
	log.Fatal(http.ListenAndServe(PORT, r))
}
