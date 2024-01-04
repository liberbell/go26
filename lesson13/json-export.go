package main

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
}
