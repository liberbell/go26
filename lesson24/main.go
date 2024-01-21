package main

import (
	"log"

	"github.com/tsawler/myniceprogram/helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	// log.Println("Hello world")

	// var myVar helpers.SomeType

	// myVar.TypeName = "some name"
	// log.Println(myVar.TypeName)

	// PrintText("Hi")make(chan int)
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
}

func PrintText(s string) {
	log.Println(s)
}
