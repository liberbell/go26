package handlers

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home_page.tpml")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about_page.tpml")
}
