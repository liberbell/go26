package main

import "os"

type Person struct {
	Name   string `json:"name"`
	Age    int    `json: "age"`
	Gender string `json: "gender"`
}

func main() {
	people := []Person{
		{"Peter", 44, "Male"},
		{"Lois", 23, "Femail"},
	}

	file, err := os.Create("export.json")
	if err != nil {
		panic(err)
	}
}
