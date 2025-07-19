package main

import "fmt"

func main() {

	// var a [5]int
	// // fmt.Println(a)

	// a[3] = 69
	// fmt.Println(a)

	// fmt.Println(len(a))

	// b := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(b)
	// fmt.Println(b[0], len(b))

	// c := []int{1, 2, 3, 5}
	// fmt.Println(len(c))

	// d := []int{1, 2, 3: 4}
	// fmt.Println(d, len(d))

	// var twod [2][3]int
	// // fmt.Println(twod)

	// for i := 0; i < 2; i++ {

	// 	for j := 0; j < 3; j++ {

	// 		twod[i][j] = i + j

	// 	}

	// }

	// fmt.Println(twod)

	e := [2][3]int{
		{0, 1, 2},
		{1, 2, 3},
	}
	fmt.Println(e)

}
