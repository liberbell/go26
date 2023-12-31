package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	aSlice := []string{"Fry", "Leela", "Bender", "Amy", "Hubert", "Zoidberg", "Hermes", "Scruffy", "Nibbler"}
	fmt.Println(aSlice)

	for i, val := range aSlice {
		fmt.Printf("%d %s\n", i, val)
	}

	aString := "One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed"
	fmt.Println(aString)

	for _, val := range aString {
		fmt.Printf("%q ", val)
	}

	aMap := map[string]string{"dog": "Scoby Doo", "cat": "Garfield", "mouse": "Jerry", "platybus": "Perry"}
	fmt.Println(aMap)

	for k, val := range aMap {
		fmt.Printf("%s: %s\n", k, val)
	}
}
