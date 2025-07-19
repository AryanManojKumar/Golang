package main

import (
	"fmt"
)

func main() {

	var i = 1

	var sum int
	for i <= 5 {

		fmt.Println(i)

		sum = sum + i
		i++
	}

	fmt.Println(sum)

}
