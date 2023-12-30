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
	fmt.Printf("%T\n", somemap2)
	fmt.Println(somemap2, len(somemap2), cap(somemap2))

	somemap2 = []string{"something#1"}
	fmt.Println(somemap2, len(somemap2), cap(somemap2))

}
