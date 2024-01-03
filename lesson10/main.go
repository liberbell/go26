package main

import (
	"fmt"
	"strings"
)

type Sidekicks interface {
	saysASaying() string
	isHuman() bool
}

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
	fmt.Println(morty.Name, "is a Rick`s poor sidekick")
	fmt.Println("Hello world")
}

func outputInfo(s Sidekicks) {
	fmt.Println("This sidekick says: ", s.saysASaying())
}
