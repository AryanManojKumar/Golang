package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"c", "a", "d"}

	slices.Sort(strs)
	fmt.Println(strs)

	ints := []int{6, 4, 1}
	slices.Sort(ints)
	fmt.Println(ints)
}
