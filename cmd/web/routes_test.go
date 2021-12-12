package main

import (
	"testing"

	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	router := routes()

	switch router.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Errorf("Router is not chi.Mux")
	}
}
