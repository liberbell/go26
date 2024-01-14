package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8000"

func main() {
	var app AppConfig

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatalln("Error creating template cache: ", err)
	}
	app.TemplateCache = tc

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting up... on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
