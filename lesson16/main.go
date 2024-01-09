package main

import "net/http"

const portNumber = "8000"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
}
