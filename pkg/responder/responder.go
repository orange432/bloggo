package responder

import (
	"encoding/json"
	"net/http"
)

type Succeeded struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type SessionResponse struct {
	Success bool   `json:"success"`
	Session string `json:"session"`
}

func Success(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Succeeded{Success: true})
}

func Session(w http.ResponseWriter, token string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SessionResponse{Success: true, Session: token})
}

func Error(w http.ResponseWriter, err string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Succeeded{Success: false, Error: err})
}
