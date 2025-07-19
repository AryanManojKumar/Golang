package main

import (
	"fmt"
)

func a() func() int {

	i := 0

	return func() int {

		i++
		return i
	}

}

func main() {

	b := a()

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

	c := a()

	fmt.Println(c())

}
