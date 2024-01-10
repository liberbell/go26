package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTemp(w http.ResponseWriter, tpml string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tpml, "./templates/base.tpml")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template: ", err)
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]
	if !inMap {
		// some cache
		log.Println("creating template and add cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}

	} else {
		log.Println("using cached template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.tpml",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = tmpl

	return nil
}
