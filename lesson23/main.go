package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Colour        string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samsong",
		Breed: "Poppy",
	}
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says: ", a.Says())
}
