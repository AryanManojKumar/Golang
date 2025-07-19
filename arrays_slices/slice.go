package main

import (
	"fmt"
)

func main() {

	// var a []string

	// a = append(a, "s")

	// fmt.Println(a, len(a) == 1)

	// b := []string{"a", "b", "c", "d", "e"}

	// c := b[0:5]

	// fmt.Println(c)

	// s := make([]string, 4)
	// // s = []string{"a", "b", "c"}

	// s[0] = "A"
	// s[1] = "B"
	// s[2] = "C"

	// // fmt.Println(s, len(s))
	// s = append(s, "E")
	// fmt.Println(s)

	// x := make([]string, len(s))
	// copy(x, s)
	// fmt.Println(x)

	// // y := s[0:3]
	// fmt.Println(s[0:3])
	// fmt.Println(s[:4])
	// fmt.Println(s[3:])

	// t := []string{"a", "b", "c"}

	// t2 := []string{"a", "b", "c"}

	// if slices.Equal(t, t2) {
	// 	fmt.Println(t, "and", t2, "are equal")
	// }

	twod := make([][]int, 3)

	for i := range 3 {

		a := i + 1
		twod[i] = make([]int, a)
		for j := range a {
			twod[i][j] = i + j
		}

		fmt.Println(twod)

	}

}
