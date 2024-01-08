package main

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home_page.tpml")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about_page.tpml")
}
