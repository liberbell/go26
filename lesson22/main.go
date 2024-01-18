package main

import "log"

func main() {
	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	// animals := []string{"dog", "cat", "horse", "pig", "cow"}
	animals := make(map[string]string)
	animals["dog"] = "Fido"
	animals["cat"] = "Fluffy"

	var firstLine = "Once upon a midnight dreary"

	for i, l := range firstLine {
		log.Println(i, ":", l)
	}
}
