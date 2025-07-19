package main

import "fmt"

func main() {

	// // a := [] make(int, 8)
	// a := []int{1, 2, 3, 4, 5, 6, 7}
	// var b int
	// var c int

	// for i := range len(a) {

	// 	if a[i]%2 == 0 {

	// 		b++

	// 	} else {
	// 		c++
	// 	}

	// }

	// fmt.Println("there are ", b, " even and  odd", c)

	// a := []int{6, 7, 8, 9}
	// // b := make([]int, len(a))
	// var y int
	// for i := 0; i <= 1; i++ {

	// 	y = a[i]
	// 	a[i] = a[3-i]
	// 	a[3-i] = y
	// }

	// fmt.Println(a)

	b := make([]int, 1)

	b = append(b, 1)
	fmt.Println(b)

}
