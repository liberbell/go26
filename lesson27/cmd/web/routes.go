package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/tsawler/go-course/pkg/handlers"

	"github.com/tsawler/go-course/pkg/config"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.Newrouter()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
