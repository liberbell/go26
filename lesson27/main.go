package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8000"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is my home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)
	_, _ = fmt.Fprintf(w, "This is about page and 2+2 is %d", sum)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValue(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
	}
}

func AddValues(x, y int) int {
	var sum int
	sum = x + y
	return sum
}

func divideValue(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
