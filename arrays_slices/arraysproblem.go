package main

import (
	"fmt"
)

func main() {

	// a := [5]int{1, 2, 3, 4, 5}
	// var b int

	// for i := 0; i <= 4; i++ {
	// 	b = b + a[i]

	// }

	// fmt.Println(b)

	var b int = 60

	a := []int{115, 20, 54, 67, 89, 60}

	c := 0

	for i := 0; i < len(a); i++ {

		if a[i] == b {

			c = 1
			fmt.Println("found ur iteam its inside the ", i, "th index")
		}

	}

	if c == 1 {

		fmt.Println("its inside array")

	} else {
		fmt.Println("its not inside array")
	}

}
