package main

import "log"

var s = "seven"

func main() {
	var s2 = "six"

	log.Println("s is: ", s)
	log.Println("s2 is: ", s2)

	saySomething("xxx")
}

func saySomething(s string) (string, string) {
	log.Println("s from the saySomething function: ", s)
	return s, "world"
}
