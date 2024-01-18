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
	firstLine = "x"

	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{})

	for i, l := range firstLine {

	}
}
