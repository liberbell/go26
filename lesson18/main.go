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
	var s2 = "six"

	s := "eight"

	log.Println("s is: ", s)
	log.Println("s2 is: ", s2)

	saySomething("xxx")
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething function: ", s)
	return s3, "world"
}
