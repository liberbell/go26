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

	whatWasSaid := saySomething()
	fmt.Println(whatWasSaid)
}

func saySomething() string {
	return "something"
}
