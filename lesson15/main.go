package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8000"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting up... on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
