package main

import (
	"fmt"

	"github.com/jagottsicher/myGoExample/helpers"
)

func main() {
	var myVar helpers.SomeType
	myVar.Name = "Stan Smith"
	fmt.Println(myVar)
}
