package middlewares

import (
	"bloggo/pkg/auth"
	"bloggo/pkg/models"
	"bloggo/pkg/render"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func VerifyToken(w http.ResponseWriter, r *http.Request) (int, error) {
	webToken, err := r.Cookie("jwt")
	if err != nil {
		render.RenderTemplate(w, "unauthorized.html", &models.TemplateData{})
		return -1, nil
	}

	// Check if token is valid
	claims := &auth.JWTClaims{}
	_, err = jwt.ParseWithClaims(
		webToken.Value,
		claims,
		func(jwtoken *jwt.Token) (interface{}, error) {
			return []byte(auth.JWT_SECRET), nil
		},
	)
	if err != nil {
		render.RenderTemplate(w, "unauthorized.html", &models.TemplateData{})
		return -1, nil
	}

	return claims.UserId, err
}

func VerifyAPIToken(w http.ResponseWriter, r *http.Request) (int, error) {
	webToken, err := r.Cookie("jwt")
	if err != nil {
		render.RenderTemplate(w, "unauthorized.html", &models.TemplateData{})
		return -1, nil
	}
	// Check if token is valid
	claims := &auth.JWTClaims{}
	_, err = jwt.ParseWithClaims(
		webToken.Value,
		claims,
		func(jwtoken *jwt.Token) (interface{}, error) {
			return []byte(auth.JWT_SECRET), nil
		},
	)

	if err != nil {
		render.RenderTemplate(w, "unauthorized.html", &models.TemplateData{})
		return -1, nil
	}

	return claims.UserId, err
}
