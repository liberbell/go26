package main

import "log"

func main() {
	var myString string
	var myInt int

	myString = "Hi"
	myInt = 21

	mySecondString := "another string"
	log.Println(myString, mySecondString, myInt)

	// var myOtherMap map[string]string

	myMap := make(map[string]string)
	myMap["dog"] = "Samson"

	log.Println(myMap["dog"])
}
