package main

import (
	"fmt"
	"time"
)

type FamilyMember struct {
	FamilyName string
	FirstName  string
	Age        int
	Birthdate  time.Time
	Spices     string
}

// var FamilyName string
// var firstName string
// var age int
// var birthdate time.Time
// var spices string

func main() {
	fmt.Println("Hello, world")

	francine := FamilyMember{
		FamilyName: "Smith",
		FirstName:  "Francine",
		Spices:     "Human",
	}
	fmt.Println(francine)

	roger := FamilyMember{
		FamilyName: "Smith",
		FirstName:  "Roger",
		Age:        1200,
		Spices:     "Alien",
	}

	fmt.Println(roger)
}
