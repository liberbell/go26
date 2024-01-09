package main

import (
	"fmt"
	"net/http"

	_ "/pkg/handlers"
)

const portNumber = ":8000"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting up... on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
