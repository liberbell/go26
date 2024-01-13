package main

import (
	"fmt"
	"net/http"

	"./pkg/config/config.go"
)

const portNumber = ":8000"

func main() {
	var app config.AppConfig

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting up... on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
