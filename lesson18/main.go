package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Trevor",
		LastName:    "Sweler",
		PhoneNumber: "1 555 5555",
	}

	log.Println(user.FirstName, user.LastName, "Birth date: ", user.BirthDate)
}

func whatever() {

}
