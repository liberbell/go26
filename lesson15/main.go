package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page.")
}

func About(w http.ResponseWriter, r *http.Request) {
	owner, saying := getData()
	sum := addValues(2 + 2)
	fmt.Fprintf(w, "This is the about page of %s \n I like to say %s \n and note is 2 + 2 = %d", owner, saying, sum)
}

func getData() (string, string) {
	o := "Eddie Murfee"
	s := "New York"
	return o, s
}

func addValues(x, y int) int {
	return x + y
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
