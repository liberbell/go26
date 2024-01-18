package main

import "log"

func main() {
	var isTrue bool

	isTrue = false
	if isTrue == true {
		log.Println("isTrue is: ", isTrue)
	} else {
		log.Println("isTrue is: ", isTrue)
	}

	cat := "dog"

	if cat == "cat" {
		log.Println("cat is: ", cat)
	} else {
		log.Println("cat is not cat")
	}

	myNum := 100
	isTrue = false

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 100 and isTrue is set to true")
	} else if myNum < 100 && isTrue {

	} else if myNum == 101 || isTrue {

	} else if myNum > 1000 && isTrue == false {

	}
}
