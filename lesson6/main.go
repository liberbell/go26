package main

import "fmt"

func main() {
	users := make(map[string]string)
	fmt.Println(users)

	users["prefer@example.com"] = "prefer example"
	fmt.Println(users)
}
