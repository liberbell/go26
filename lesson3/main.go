package main

import "fmt"

var role1 = "Billionere"

func main() {
	var role2 = "Inventor"

	fmt.Println("Bruce wayne is ", role1)
	fmt.Println("Bruce wayne is ", role2)

	tellmeWhoYouAre("Batman")
}

func tellmeWhoYouAre(role3 string) string {
	fmt.Println("Bruce wayne is ", role1)
	return role3
}
