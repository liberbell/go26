package main

import (
	"log"
	"time"
)

var s = "seven"

func main() {
	var s2 = "six"

	s := "eight"
	var firstName string
	var lastName string
	var phoneNumber string
	var age int
	var birthDate time.Time

	log.Println("s is: ", s)
	log.Println("s2 is: ", s2)

	saySomething("xxx")
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething function: ", s)
	return s3, "world"
}
