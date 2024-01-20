package main

type Animal interface {
	Say() string
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
