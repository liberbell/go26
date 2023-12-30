package main

import "fmt"

func main() {
	users := make(map[string]string)
	fmt.Println(users)
	fmt.Println(len((users)))

	users["prefer@example.com"] = "prefer example"
	users["lois@example.com"] = "lois example"
	users["meg@example.com"] = "meg example"
	users["chris@example.com"] = "chris example"
	users["stewie@example.com"] = "stewie example"
	users["brian@example.com"] = "brian example"

	fmt.Println(users["meg@example.com"])
	fmt.Println(users)

	users["stewie@example.com"] = "stewerd example"
	fmt.Println(users)
	fmt.Println(len(users))

	delete(users, "stewie@example.com")
	fmt.Println(users)
	fmt.Println(len(users))
}
