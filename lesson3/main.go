package main

import "fmt"

var role1 = "Billionere"

func main() {
	role1 := "Playboy"
	var role2 = "Inventor"

	fmt.Println("Bruce wayne is ", role1)
	fmt.Println("Bruce wayne is ", role2)

	tellmeWhoYouAre("Batman")
}

func tellmeWhoYouAre(role1 string) string {
	fmt.Println("Bruce wayne is ", role1)
	return role1
}
