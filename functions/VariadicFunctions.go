package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Println(nums)

	var total int

	for _, num := range nums {

		total = total + num

	}

	fmt.Println(total)
}

func main() {

	sum(1, 2, 3, 4, 5)

	a := []int{10, 20, 30, 40}
	sum(a...)
}
