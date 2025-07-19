package main

import (
	"fmt"
)

func main() {

	// nums := []int{7, 8, 9, 4, 5, 6}
	// var b int

	// for _, i := range nums {

	// 	b = b + i

	// }

	// fmt.Println(b)

	// for i, num := range nums {
	// 	if num == 4 {
	// 		fmt.Println(i + 1)
	// 	}
	// }

	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cabage"}

	for k, v := range kvs {
		fmt.Println(k, "->", v)
	}
}
