package main

import "os"

type Person struct {
	Name   string `json:"name"`
	Age    int    `json: "age"`
	Gender string `json: "gender"`
}

func main() {
	file, err := os.Open("data.json")
}
