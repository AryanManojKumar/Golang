package main

import (
	"fmt"
)

// func val() (int, int) {
// 	return 5, 5
// }

// func main() {

// 	b, a := val()
// 	var c int
// 	c = a + b

// 	fmt.Println(c)

func cal(a, b int) (int, int, int, int) {

	// sum := a + b
	// div := a / b
	// multi := a * b
	// sub := a - b

	// fmt.Println(sum, div, multi, sub)
	// return 0

	return a + b, a - b, a / b, a * b
}

func main() {

	var d, e int
	fmt.Println("enter two numbers to perform multiple arithmetic operations")
	fmt.Scan(&d, &e)
	sum, sub, div, multi := cal(d, e)
	fmt.Println(sum, sub, div, multi)

}
