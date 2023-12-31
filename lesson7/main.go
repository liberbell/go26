package main

import "fmt"

func main() {
	var somemap [3]string
	fmt.Println(somemap)
	fmt.Println(len(somemap))

	somemap[0] = "something#1"
	somemap[1] = "something#2"
	somemap[2] = "something#3"
	fmt.Println(somemap)

	// somemap[3] = "something#4"

	var somemap2 []string
	fmt.Printf("%T\n", somemap2)
	fmt.Println(somemap2, len(somemap2), cap(somemap2))

	somemap2 = []string{"something#1"}
	fmt.Println(somemap2, len(somemap2), cap(somemap2))

	someInteger := []int{42, 23, 1000, -1}
	fmt.Printf("%T\n", someInteger)
	fmt.Println(someInteger, len(someInteger), cap(someInteger))
	someInteger = append(someInteger, 987, 654, 765, 432, 321, 543)
	fmt.Println(someInteger, len(someInteger), cap(someInteger))
	someInteger = append(someInteger, 999)
	fmt.Println(someInteger, len(someInteger), cap(someInteger))

	someFloat := make([]float64, 3, 6)
	fmt.Printf("%T\n", someFloat)
	fmt.Println(someFloat, len(someFloat), cap(someFloat))

	someFloat[0] = 42.42
	fmt.Println(someFloat, len(someFloat), cap(someFloat))
	someFloat = append(someFloat, 41.41, 445.998)
	fmt.Println(someFloat, len(someFloat), cap(someFloat))
	someFloat = append(someFloat, 666.666, 777.777)

}
