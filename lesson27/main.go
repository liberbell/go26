package main

import (
	"fmt"
	"net/http"
)

var portNumber = ":8000"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello world")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes: %d", n))
	})

	http.ListenAndServe(portNumber, nil)
}
