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
	} else {
		fmt.Println(number, "is zero or positive")
	}
}
