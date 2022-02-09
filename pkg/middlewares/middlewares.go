package middlewares

import (
	"bloggo/pkg/auth"
	"bloggo/pkg/models"
	"bloggo/pkg/render"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func VerifyToken(w http.ResponseWriter, r *http.Request) error {
	webToken, _ := r.Cookie("jwt")
	if webToken.Value == "" {

		render.RenderTemplate(w, "unauthorized.html", &models.TemplateData{})
	}
	// Check if token is valid
	_, err := jwt.ParseWithClaims(
		webToken.Value,
		&auth.JWTClaims{},
		func(jwtoken *jwt.Token) (interface{}, error) {
			return []byte(auth.JWT_SECRET), nil
		},
	)
	fmt.Println(err)
	return err
}
