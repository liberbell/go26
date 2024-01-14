package main

import (
	"fmt"
)

func main() {
	var whatToSay string
	var i int

	whatToSay = "Goodbye cruel world"
	fmt.Println(whatToSay)

	i = 8
	fmt.Println("i is set to: ", i)

	saySomething()
}

func saySomething() string {
	return "something"
}
