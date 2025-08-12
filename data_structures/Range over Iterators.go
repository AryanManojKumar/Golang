package main

import (
	"fmt"
	"iter"
)

// a := []int {1,2,3}

// func counttothree() iter.Seq[int] {
// 	return func(yield func(int) bool) {
// 		for i := 0; i < 3; i++ {
// 			if !yield(i){
// 				return
// 			}

// 		}
// 	}
// }

func geneven(max int) iter.Seq[int] {
	return func(yield func(int) bool) {

		for i := 0; i < max; i++ {

			if i%2 == 0 {

				if !yield(i) {
					return
				}
			}
		}
	}
}

func main() {

	for n := range geneven(10) {
		fmt.Println(n)
	}
}
