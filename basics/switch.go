package main

import (
	"fmt"
)

func main() {

	// switch time.Now().Weekday() {

	// case time.Saturday, time.Sunday:
	// 	fmt.Println("its weekend")

	// default:
	// 	fmt.Println("its week-days")
	// }

	// i := 1
	// fmt.Println("write", i, "as")
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("idk how to write")
	// }

	// fmt.Println("today is ")

	// a := time.Now().Hour()
	// switch {
	// case a > 12:
	// 	fmt.Println("its pass 12")
	// default:
	// 	fmt.Println("its not pass 12")
	// }

	// fmt.Println(a)

	// switch time.Now().Month() {

	// case time.June, time.May:
	// 	fmt.Println("birthday month")

	// case time.February, time.July:
	// 	fmt.Println("bad month")
	// }

	x := 10
	y := 10

	op := "/"

	switch op {

	case "+":
		fmt.Println(x + y)
	case "*":
		fmt.Println(x * y)

	case "/":
		if y > 0 {
			fmt.Println(x / y)
		} else if y < 0 {
			fmt.Println("can not divide negative ")
		} else {
			fmt.Println("can not divide zero")
		}
	}

}
