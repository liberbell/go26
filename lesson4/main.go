package main

import (
	"fmt"
	"time"
)

type FamilyMember struct {
	familyName string
	firstName  string
	age        int
	birthdate  time.Time
	spices     string
}

// var FamilyName string
// var firstName string
// var age int
// var birthdate time.Time
// var spices string

func main() {
	fmt.Println("Hello, world")

	francine := FamilyMember{
		familyName: "Smith",
		firstName:  "Francine",
		spices:     "Human",
	}
	fmt.Println(francine)
}
