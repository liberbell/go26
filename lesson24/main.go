package main

import (
	"command-line-arguments/Users/hideakiehara/virtualenv/go26/lesson24/helpers/helpers.go"
	"log"
)

func main() {
	log.Println("Hello world")

	var myVar helpers.SomeType

	myVar.TypeName = "some name"
	log.Println(myVar.TypeName)
}
