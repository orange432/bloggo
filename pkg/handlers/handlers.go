package handlers

import (
	"bloggo/pkg/models"
	"bloggo/pkg/render"
	"encoding/json"
	"log"
	"net/http"
)

type Success struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type Login struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Session string `json:"session"`
}

func Home(w http.ResponseWriter, r *http.Request) {

	testMap := make(map[string]string)
	testMap["test"] = "Daniel is testing this page!"

	render.RenderTemplate(w, "home.html", &models.TemplateData{
		StringMap: testMap,
	})
}

func TestPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "test.html", &models.TemplateData{})
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html", &models.TemplateData{})
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login.html", &models.TemplateData{})
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "register.html", &models.TemplateData{})
}

func Register(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Success{Success: true})
}

func SaveArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Success{Success: true})
	if err != nil {
		log.Fatalf("error while encoding JSON")
	}
}
