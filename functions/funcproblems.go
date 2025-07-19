package main

import (
	"fmt"
)

func check(vows ...string) {

	vowels := []string{"a", "e", "i", "o", "u"}
	var a int

	for j := range vowels {

		for _, vow := range vows {
			if vow == vowels[j] {
				fmt.Println("it has vowels ", vowels[j])
				a++

			}
		}
	}
	fmt.Println(a)

}

func main() {
	check("f", "d", "e", "m")
}
