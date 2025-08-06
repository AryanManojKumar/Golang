package main

import (
	"fmt"
)

func main() {

	a := 0
	for a < 5 {
		fmt.Println(a)
		a++
	}

	for i := 0; i < 100; i++ {

		fmt.Println(i)

	}

	for {
		fmt.Println("loop runing")
		break

	}

	for b := 0; b < 100; b++ {

		if b%2 == 0 {
			continue
		}
		fmt.Println(b)

	}

}
