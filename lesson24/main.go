package main

import (
	"log"

	"github.com/tsawler/myniceprogram/helpers"
)

func main() {
	log.Println("Hello world")

	var myVar helpers.SomeType

	myVar.TypeName = "some name"
	log.Println(myVar.TypeName)
}
