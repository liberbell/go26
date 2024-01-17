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
	myMap["other-dog"] = "Cassie"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])
}
