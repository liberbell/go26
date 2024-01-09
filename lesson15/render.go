package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tpml string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tpml, "./templates/base.tpml")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template: ", err)
	}
}

var tc = make(map[string])

func renderTemplateTemp(w http.ResponseWriter, t string)  {
	
}