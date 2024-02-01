package handlers

import (
	"net/http"

	"github.com/tsawler/go-course/pkg/config"
	"github.com/tsawler/go-course/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home_page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about_page.tmpl")
}
