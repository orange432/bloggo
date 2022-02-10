package middlewares

import (
	"bloggo/pkg/auth"
	"bloggo/pkg/models"
	"bloggo/pkg/render"
	"bloggo/pkg/responder"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(w http.ResponseWriter, r *http.Request) (uint, error) {
	webToken, _ := r.Cookie("jwt")
	if webToken.Value == "" {

		render.RenderTemplate(w, "unauthorized.html", &models.TemplateData{})
	}
	// Check if token is valid

	claims := &auth.JWTClaims{}
	_, err := jwt.ParseWithClaims(
		webToken.Value,
		claims,
		func(jwtoken *jwt.Token) (interface{}, error) {
			return []byte(auth.JWT_SECRET), nil
		},
	)

	return claims.UserId, err
}

func VerifyAPIToken(w http.ResponseWriter, r *http.Request) (uint, error) {
	webToken, _ := r.Cookie("jwt")
	if webToken.Value == "" {
		responder.Error(w, "Please log in.")
	}
	// Check if token is valid
	claims := &auth.JWTClaims{}
	_, err := jwt.ParseWithClaims(
		webToken.Value,
		claims,
		func(jwtoken *jwt.Token) (interface{}, error) {
			return []byte(auth.JWT_SECRET), nil
		},
	)
	return claims.UserId, err
}
