package main

import (
	"fmt"
	"net/http"

	"github.com/tsawler/go-course/pkg/handlers"
)

const portNumber = ":8000"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
