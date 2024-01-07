package main

import (
	"errors"
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
	f, err := divideValue(x, y)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error: division by zero is not a valid operation. Error returned: %s", err)
		return
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
}

func divideValue(x, y float64) (float64, error) {
	var divide float64
	if y == 0 {
		err := errors.New("divided by zero")
		return 0, err
	}
	divide = x / y
	return divide, nil
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
