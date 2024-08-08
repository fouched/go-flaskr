package main

import (
	"github.com/fouched/go-flaskr/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes() http.Handler {

	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	mux.Get("/register", handlers.RegisterGet)
	mux.Post("/register", handlers.RegisterPost)
	return mux
}
