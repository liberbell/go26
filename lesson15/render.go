package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tpml string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tpml, "./templates/base.tpml")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template: ", err)
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplateTemp(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]
	if !inMap {
		// some cache
	} else {
		// no make
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
