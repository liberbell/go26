package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tsawler/go-course/pkg/config"
	"github.com/tsawler/go-course/pkg/handlers"
	"github.com/tsawler/go-course/pkg/render"
)

const portNumber = ":8000"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

// 	_ = http.ListenAndServe(portNumber, nil)
// }
