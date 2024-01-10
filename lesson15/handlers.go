package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplateTemp(w, "home_page.tpml")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplateTemp(w, "about_page.tpml")
}
