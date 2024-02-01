package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tsawler/go-course/pkg/config"

	"github.com/tsawler/go-course/pkg/handlers"
)

const portNumber = ":8000"

func main() {

	var app config.AppConfig

	tc, err := render.createTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
