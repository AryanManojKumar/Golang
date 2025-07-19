package main

import (
	"fmt"
)

// func a(zeroval int) {
// 	zeroval = 0
// }

// func b(zeroptr *int) {
// 	*zeroptr = 0
// }

// func main() {

// i := 1
// fmt.Println("staring:", i)

// a(i)
// fmt.Println("after call by value ", i)
// b(&i)
// fmt.Println("after call by refrence ", i)
// fmt.Println("pointer", &i)

// a := 5
// b := &a

// fmt.Println("value of a ", a)
// fmt.Println("address of a ", b)
// fmt.Println("what value is the b pointing to", *b)

// }

func multi(value *int) {
	*value = *value * 10
}

func main() {
	a := 50

	fmt.Println("initial a value", a)
	multi(&a)
	fmt.Println("after func a value ", a)
}
