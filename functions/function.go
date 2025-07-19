package main

import (
	"fmt"
)

func plus(a, b float32) float32 {

	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func main() {

	a := plus(5.01, 5.05)
	fmt.Println(a)
}
