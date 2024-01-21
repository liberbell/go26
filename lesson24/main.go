package main

import (
	"command-line-arguments/Users/hideakiehara/Go/src/github.com/tsawler/myniceprogram/helpers/helpers.go"
	"log"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
}

func main() {
	// log.Println("Hello world")

	// var myVar helpers.SomeType

	// myVar.TypeName = "some name"
	// log.Println(myVar.TypeName)

	PrintText("Hi")make(chan int)
	intChan := 
}

func PrintText(s string) {
	log.Println(s)
}
