package main

type People struct {
	FirstName string `json: "first_name"`
	LastName  string `json: "last_name"`
	HairClour string `json: "hair_colur"`
	HasDog    bool   `json: "has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "John",
			"last_name": "Denver",
			"hair_colour": "brown",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_colour": "black",
			"has_dog": false
		}
	]`

}
