package main

import (
	"fmt"
)

func main() {

	if 6%2 == 0 {

		fmt.Println("its even")
	} else {
		fmt.Println("its odd")
	}

	if 12%2 == 0 {
		fmt.Println("12 is divided by 2")
	}

	if num := 1; num < 0 {
		fmt.Println("its  a negative number")

	} else if num > 9 {
		fmt.Println("it has more then 1 digit")
	} else {
		fmt.Println("it has single digit")
	}

}
