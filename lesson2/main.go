package main

import "fmt"

func main() {
	aCostume := "Clip Clop"
	fmt.Printf("Roger is wearing a costume1 as %s.\n", aCostume)

	changeTheCostume(aCostume)
	fmt.Printf("Roger is wearing a costume2 as %s.\n", aCostume)
}

func changeTheCostume(c string) {
	fmt.Printf("Roger is wearing a costume3 as %s.\n", c)
	newCostume := "Rickey Spanish"

	c = newCostume
	fmt.Printf("Roger is wearing a costume4 as %s.\n", c)
}
