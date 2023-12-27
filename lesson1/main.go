package main

import "fmt"

var x = 23

func main() {
	fmt.Println("Hello, world")
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)
	y := 42.0
	fmt.Println(y)
	fmt.Printf("y is of type %T\n", y)
	z := "goose"
	fmt.Println(z)
	fmt.Printf("z is of type %T\n", z)

	done, trueOrNot := outputThat(x, y, z)
	fmt.Println(done, trueOrNot)
}

func outputThat(a int, b float64, c string) (string, bool) {
	fmt.Printf("Integer: %d, FloatinPoint: %f, String: %s\n", a, b, c)
	return "Ok,", true
}
