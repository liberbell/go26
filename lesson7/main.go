package main

import "fmt"

func main() {
	var somemap [3]string
	fmt.Println(somemap)
	fmt.Println(len(somemap))

	somemap[0] = "something#1"
	somemap[1] = "something#2"
	somemap[2] = "something#3"
	fmt.Println(somemap)

	// somemap[3] = "something#4"

	var somemap2 []string

}
