package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8000"

func Home(w http.ResponseWriter, r *http.Request) {

}

func About(w http.ResponseWriter, r *http.Request) {

}

func renderTemplate(w http.ResponseWriter, tpml string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tpml)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template: ", err)
	}
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting up... on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
