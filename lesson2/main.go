package main

import "fmt"

func main() {
	aCostume := "Clip Clop"
	fmt.Printf("Roger is wearing a costume as %s.\n", aCostume)

	changeTheCostume(aCostume)
	fmt.Printf("Roger is wearing a costume as %s.\n", aCostume)
}

func changeTheCostume(c string) {
	fmt.Printf("Roger is wearing a costume as %s.\n", c)
	newCostume := "Rickey Spanish"

}
