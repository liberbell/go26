package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8000"

func Home(w http.ResponseWriter, r *http.Request) {

}

func About(w http.ResponseWriter, r *http.Request) {

}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting up... on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
