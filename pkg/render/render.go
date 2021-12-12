package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/orange432/bloggo/pkg/config"
	"github.com/orange432/bloggo/pkg/models"
)

var app *config.AppConfig

var PATH_TO_PAGES = "./templates/pages"
var PATH_TO_LAYOUTS = "./templates/layouts"

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	var templateCache map[string]*template.Template

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	// Load individual tem-late
	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Couldn't load templates.")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, data)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", w)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.html", PATH_TO_PAGES))
	if err != nil {
		return tCache, err
	}

	for _, page := range pages {
		// Load normal page content
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tCache, err
		}

		// Load layouts
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.html", PATH_TO_LAYOUTS))
		if err != nil {
			return tCache, err
		}

		// Layouts found, add to template string
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.html", PATH_TO_LAYOUTS))
			if err != nil {
				return tCache, err
			}
		}
		tCache[name] = ts
	}
	// Gone through each page and template, return success
	return tCache, nil
}
