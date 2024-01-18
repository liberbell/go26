package main

import "log"

func main() {
	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	// animals := []string{"dog", "cat", "horse", "pig", "cow"}
	animals := make(map[string]string)

	for _, animal := range animals {
		log.Println(animal)
	}
}
