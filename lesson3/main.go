package main

import "fmt"

var role1 = "Billionere"

func main() {
	var role2 = "Inventor"

	fmt.Println("Bruce wayne is ", role1)
	fmt.Println("Bruce wayne is ", role2)
}

func tellmeWhoYouAre(role1 string) string {
	return role1
}
