package main

type FamilyMember struct {
	FamilyName string
	FirstNmae  string
	Age        int
	Species    string
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
}
