package main

import (
	"github.com/fouched/go-flaskr/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes() http.Handler {

	mux := chi.NewRouter()
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Instance.Home)
	mux.Get("/register", handlers.Instance.RegisterGet)
	mux.Post("/register", handlers.Instance.RegisterPost)
	mux.Get("/login", handlers.Instance.LoginGet)
	mux.Post("/login", handlers.Instance.LoginPost)
	mux.Get("/logout", handlers.Instance.Logout)

	return mux
}
