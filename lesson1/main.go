package main

import "fmt"

var x = 23

func main() {
	fmt.Println("Hello, world")
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)
	y := 42.0
	fmt.Println(y)
	fmt.Printf("y is of type %T\n", y)
}
