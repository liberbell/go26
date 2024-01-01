package main

import (
	"fmt"
	"log"
)

func main() {
	number := 23
	log.Println("Let`s talk about", number)

	if number < 0 {
		fmt.Println(number, "is negatie")
	} else if number < 10 {
		fmt.Println(number, "has 1 digit")
	} else {
		fmt.Println(number, "has multiple digits")
	}

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}
}
