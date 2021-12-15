package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/orange432/bloggo/pkg/config"
	"github.com/orange432/bloggo/pkg/render"
)

const PORT_NUMBER = ":4000"

func main() {

	err := run()

	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    PORT_NUMBER,
		Handler: routes(),
	}

	// Start the server
	fmt.Println("🚀 Started at http://localhost" + PORT_NUMBER)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	var app config.AppConfig

	// Load the pages
	tCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Println(err)
		log.Fatal("Can't load template cache")
		return err
	}

	app.TemplateCache = tCache

	render.NewTemplates(&app)

	return nil
}
