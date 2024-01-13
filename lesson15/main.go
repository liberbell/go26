package main

import (
	"command-line-arguments/Users/hideakiehara/virtualenv/go26/lesson15/pkg/config/config.go"
	"fmt"
	"net/http"
)

const portNumber = ":8000"

func main() {
	var app config.AppConfig

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting up... on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
