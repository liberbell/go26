package main

import (
	"log"
	"sort"
)

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

	var mySlice []string
	mySlice = append(mySlice, "Travor")
	mySlice = append(mySlice, "John")
	log.Println(mySlice)

	var mySlice2 []int
	mySlice2 = append(mySlice2, 2)
	mySlice2 = append(mySlice2, 1)
	mySlice2 = append(mySlice2, 3)
	sort.Ints(mySlice2)
	log.Println(mySlice2)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)
	log.Println(numbers[6:9])

	names := []string{"one", "two", "seven", "fish", "cat"}
	log.Println(names)

}
