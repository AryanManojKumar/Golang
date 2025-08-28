package main

import (
	"cmp"
	"fmt"
	"slices"
)

// func main() {

// 	fruits := []string{"dragonfruit", "apple", "banana"}

// 	customfunc := func(a, b string) int {
// 		return cmp.Compare(len(a), len(b))

// 	}

// 	slices.SortFunc(fruits, customfunc)
// 	fmt.Print(fruits)
// }

func main() {

	type cars struct {
		name string
		age  int
	}

	bmw := []cars{
		cars{name: "m4 competition", age: 9},
		cars{name: "m8", age: 3},
	}

	ageshort := func(a, b cars) int {

		return cmp.Compare(a.age, b.age)

	}

	slices.SortFunc(bmw, ageshort)
	fmt.Println(bmw)

}
