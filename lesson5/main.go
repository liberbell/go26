package main

import "fmt"

type FamilyMember struct {
	FamilyName string
	FirstNmae  string
	Age        int
	Species    string
}

func (r FamilyMember) sayYourName() string {
	fullName := r.FirstNmae + " " + r.FamilyName
	return fullName
}

func main() {
	francine := FamilyMember{
		FamilyName: "Smith",
		FirstNmae:  "Francie",
		Age:        41,
		Species:    "Human",
	}

	stan := FamilyMember{
		FamilyName: "Smith",
		FirstNmae:  "Stan",
		Age:        44,
		Species:    "Human",
	}

	steve := FamilyMember{
		FamilyName: "Smith",
		FirstNmae:  "Steve Anita",
		Age:        13,
		Species:    "Human",
	}

	hayley := FamilyMember{
		FamilyName: "Smith",
		FirstNmae:  "Hayley",
		Age:        19,
		Species:    "Human",
	}

	jeff := FamilyMember{
		FamilyName: "Fisher",
		FirstNmae:  "Jeff",
		Age:        31,
		Species:    "Human",
	}

	Klous := FamilyMember{
		FamilyName: "Heibler",
		FirstNmae:  "Klous",
		Age:        23,
		Species:    "Fish",
	}

	roger := FamilyMember{
		FamilyName: "Smith",
		FirstNmae:  "Roger",
		Age:        1277,
		Species:    "Alien",
	}

	fmt.Println(francine)
	fmt.Println(stan)
	fmt.Println(hayley)
	fmt.Println(jeff)
	fmt.Println(steve)
	fmt.Println(Klous)
	fmt.Println(roger)

	fmt.Println("Hello my name is ", roger.sayYourName())
}
