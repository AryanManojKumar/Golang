package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["apple"] = 5
	m["grape"] = 8

	// fmt.Println()

	// v1 := m["apple"]
	// v2 := m["lemon"]

	// fmt.Println(v2, v1, len(m))

	// delete(m, "apple")

	// clear(m)
	// fmt.Println(m)

	// _, prs := m["apple"]

	// fmt.Println(prs)

	n := map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}

	var a int

	for i := range 5 {

		a = a + n[i]

	}

	fmt.Println(a)

	f := map[string]int{"apple": 1, "gavava": 2}

	fmt.Println(f)

}
