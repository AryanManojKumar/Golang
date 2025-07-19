package main

import (
	"fmt"
)

func main() {

	grade := 49
	b := "congratulations you passed , ur grades are "

	if grade >= 50 {

		fmt.Print(b)

		if grade >= 90 {

			fmt.Println("A")

		} else if grade >= 70 && grade <= 89 {
			fmt.Println("B")
		} else if grade >= 50 && grade <= 69 {
			fmt.Println("C")
		}

	} else {
		fmt.Println("sorry you failed")
	}
}
