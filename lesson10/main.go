package main

import (
	"fmt"
	"strings"
)

type Wookiee struct {
	Name     string
	Spieces  string
	Fraction string
}

type Boy struct {
	Name string
	Age  int
}

func main() {
	chewbacca := Wookiee{
		Name:     "Chewbacca",
		Spieces:  "Wookiee",
		Fraction: "Rebellion",
	}

	morty := Boy{
		Name: "Morty",
		Age:  14,
	}

	fmt.Println(chewbacca.Name, "is a tall", strings.ToLower(chewbacca.Spieces))
	fmt.Println("Hello world")
}
