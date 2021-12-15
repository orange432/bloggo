package handlers

import (
	"encoding/json"
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/orange432/bloggo/pkg/config"
	"github.com/orange432/bloggo/pkg/models"
	"github.com/orange432/bloggo/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	testMap := make(map[string]string)
	testMap["test"] = "Daniel is testing this page!"

	render.RenderTemplate(w, "home.html", &models.TemplateData{
		StringMap: testMap,
	})
}

func (m *Repository) TestPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "test.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html", &models.TemplateData{})
}

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login.html", &models.TemplateData{})
}

type LoginDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (m *Repository) LoginPost(w http.ResponseWriter, r *http.Request) {
	var login LoginDetails
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&login)
	if err != nil {
		log.Println(err)
		resMap := make(map[string]string)
		resMap["error"] = "Error loading request."

		w.Header().Set("Content-Type", "application/json")
		jsonRes, _ := json.Marshal(resMap)
		w.Write(jsonRes)
		return
	}

	log.Println(login.Username)
	resMap := make(map[string]string)
	resMap["message"] = "Login Successful!"
	resMap["success"] = "true"

	w.Header().Set("Content-Type", "application/json")
	jsonRes, _ := json.Marshal(resMap)
	w.Write(jsonRes)
}

func (m *Repository) Register(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "register.html", &models.TemplateData{})
}

func (m *Repository) RegisterPost(w http.ResponseWriter, r *http.Request) {
	var login LoginDetails
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&login)
	if err != nil {
		log.Println(err)
		resMap := make(map[string]string)
		resMap["error"] = "Error loading request."
		resMap["success"] = "false"

		w.Header().Set("Content-Type", "application/json")
		jsonRes, _ := json.Marshal(resMap)
		w.Write(jsonRes)
		return
	}

	log.Println(login.Username)
	resMap := make(map[string]string)
	resMap["message"] = "Registration Successful!"
	resMap["success"] = "true"

	w.Header().Set("Content-Type", "application/json")
	jsonRes, _ := json.Marshal(resMap)
	w.Write(jsonRes)
}

func (m *Repository) GetArticle(w http.ResponseWriter, r *http.Request) {
	articleSlug := chi.URLParam(r, "articleSlug")
	articlePath := "./articles/" + articleSlug + ".html"
	var articleContent string

	resMap := make(map[string]string)
	resMap["title"] = articleSlug

	_, err := os.Stat(articlePath)

	if errors.Is(err, os.ErrNotExist) {
		articleContent = "Article not found!"
	} else {
		content, err := ioutil.ReadFile(articlePath)
		if err != nil {
			log.Println(err)
		}
		articleContent = string(content[:])
	}

	render.RenderTemplate(w, "article.html", &models.TemplateData{
		StringMap: resMap,
		Article:   template.HTML(articleContent),
	})
}
