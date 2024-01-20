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
	users = append(users, User{"John", "Smith", "john.smith@example.com", 20})
	users = append(users, User{"Mary", "Jones", "mary.jones@example.com", 56})
	users = append(users, User{"Sally", "Brown", "sally.brown@example.com", 71})
	users = append(users, User{"Alex", "Anderson", "alex.anderson@example.com", 4})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
