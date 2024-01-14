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

	whatWasSaid, theOtherThing := saySomething()
	fmt.Println("The function returned: ", whatWasSaid, theOtherThing)
}

func saySomething() (string, string) {
	return "something", "returned"
}
