package responder

import (
	"encoding/json"
	"net/http"
	"time"
)

type Succeeded struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func Success(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Succeeded{Success: true})
}

func Session(w http.ResponseWriter, token string) {
	expiration := time.Now().Add(12 * time.Hour)
	cookie := http.Cookie{Name: "jwt", Value: token, Expires: expiration, Path: "/", HttpOnly: true, Secure: true}
	http.SetCookie(w, &cookie)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Succeeded{Success: true})
}

func Error(w http.ResponseWriter, err string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Succeeded{Success: false, Error: err})
}
