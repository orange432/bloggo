package handlers

import (
	"bloggo/pkg/models"
	"bloggo/pkg/render"
	"bloggo/pkg/responder"
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

type LoginDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTClaims struct {
	UserId   uint   `json: "userid"`
	Username string `json: "username"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	var l LoginDetails

	// Get the request body
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&l)
	if err != nil {
		responder.Error(w, "Invalid login data.")
		return
	}

	db, err := gorm.Open(sqlite.Open("bloggo.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		responder.Error(w, "Failed to connect to database")
		return
	}
	db.AutoMigrate(&UserModel{})

	// Get user
	var user UserModel
	result := db.First(&user, "Username = ?", l.Username)

	// User wasn't found
	if result.RowsAffected == 0 {
		responder.Error(w, "User does not exist")
		return
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(l.Password))
	if err != nil {
		responder.Error(w, "Invalid Username/Password.")
		return
	}

	// Password and username is valid, create token
	claims := JWTClaims{
		UserId:   user.UserId,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "Bloggo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("SECRET"))

	responder.Session(w, signedToken)
}

type UserModel struct {
	gorm.Model
	UserId      uint   `gorm:"primaryKey;autoIncrement"`
	Username    string `gorm:"unique"`
	Password    string
	DisplayName string
}

type UserDetails struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	DisplayName string `json:"displayname"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var u UserDetails

	// Get the request body
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&u)
	if err != nil {
		responder.Error(w, "Invalid registration data.")
		return
	}

	db, err := gorm.Open(sqlite.Open("bloggo.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		responder.Error(w, "Failed to connect to database")
		return
	}
	db.AutoMigrate(&UserModel{})

	// Check if user exists
	var user UserModel
	result := db.First(&user, "Username = ?", u.Username)

	if result.RowsAffected != 0 {
		responder.Error(w, "Username already exists.")
		return
	}

	// Hash password
	passwordBytes := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		responder.Error(w, "Error while hashing password.")
		return
	}

	details := UserModel{
		Username:    u.Username,
		Password:    string(hashedPassword),
		DisplayName: u.DisplayName,
	}

	res := db.Create(&details)

	if res.Error != nil {
		log.Println(res.Error)
		responder.Error(w, "Error adding data to the database")
		return
	}
	responder.Success(w)
}
