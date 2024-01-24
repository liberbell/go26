package main

import (
	"fmt"
	"net/http"
)

var portNumber = ":8000"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is my home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)
	fmt.Fprintf(w, "This is about page and 2+2 is %d", sum)
}

func AddValues(x, y int) int {
	var sum int
	sum = x + y
	return sum
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(portNumber, nil)
}
