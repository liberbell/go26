package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page.")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {})

	http.ListenAndServe("localhost:8000", nil)
}
