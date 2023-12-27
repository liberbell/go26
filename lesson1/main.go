package main

import "fmt"

var x = 23

func main() {
	fmt.Println("Hello, world")
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)
	y := 4.2
	fmt.Printf("y is of type %T\n", y)
}
