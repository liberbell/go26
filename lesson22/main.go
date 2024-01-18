package main

import "log"

func main() {
	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "cat", "horse", "pig", "cow"}

	for i, animal := range animals {
		log.Println(i, animal)
	}
}
