package main

import "log"

func main() {
	var myString string
	myString = "green"

	log.Println("myString is set to : ", myString)
	changeUsingPointer(&myString)
}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
