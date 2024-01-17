package main

import "log"

// type User struct {
// 	FirstName string
// 	LastName  string
// }

func main() {
	var myString string
	var myInt int

	myString = "Hi"
	myInt = 21

	mySecondString := "another string"
	log.Println(myString, mySecondString, myInt)

	// var myOtherMap map[string]string

	// myMap := make(map[string]string)
	// myMap["dog"] = "Samson"
	// myMap["other-dog"] = "Cassie"
	// myMap["dog"] = "fido"

	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])

	// myMap := make(map[string]int)

	// myMap["first"] = 1
	// myMap["second"] = 2

	// log.Println(myMap["first"])
	// log.Println(myMap["second"])

	// myMap := make(map[string]User)

	// me := User{
	// 	FirstName: "Trevor",
	// 	LastName:  "Sawler",
	// }

	// myMap["me"] = me
	// log.Println(myMap["me"].FirstName)

	// var myNewVar float64
	// myNewVar = 11.1

	var myString string
	myString = "Fish"
	log.Println(myString)

}
