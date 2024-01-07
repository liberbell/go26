package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8000"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page.")
}

func About(w http.ResponseWriter, r *http.Request) {
	owner, saying := getData()
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, "This is the about page of %s \n I like to say %s \n and note is 2 + 2 = %d", owner, saying, sum)
}

func getData() (string, string) {
	o := "Eddie Murfee"
	s := "New York"
	return o, s
}

func addValues(x, y int) int {
	var sum int
	sum = x + y
	return sum
}

func Divide(w http.ResponseWriter, r *http.Request) {
	x := 100.0
	y := 10.0
	divideValue(x, y)
}

func divideValue(x, y float64) float64 {
	var divide float64
	divide = x / y
	return divide
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello, world")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	// })

	fmt.Println(fmt.Sprintf("Starting up... on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
