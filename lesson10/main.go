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

func (r *Wookiee) saysASaying() string {
	return "\"RAWRGWAGGR!\""
}

func (r *Wookiee) isHuman() bool {
	return false
}

type Boy struct {
	Name string
	Age  int
}

func (r *Boy) saysASaying() string {
	return "\"Why I don`t have a chachphrase?\""
}

func (r *Boy) isHuman() bool {
	return true
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
	outputInfo(&chewbacca)
	fmt.Println(morty.Name, "is a Rick`s poor sidekick")
	outputInfo(&morty)
	// fmt.Println("Hello world")
}

func outputInfo(s Sidekicks) {
	fmt.Println("This sidekick says: ", s.saysASaying())
	if s.isHuman() == true {
		fmt.Println("and is human")
	} else {
		fmt.Println("and is not human")
	}
	fmt.Println()
}
