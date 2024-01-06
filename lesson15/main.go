package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page.")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

func getData() (string, string) {
	o := "Eddie Murfee"
	s := "New York"
	return o, s
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello, world")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	// })

	http.ListenAndServe("localhost:8000", nil)
}
