package main

import (
	"errors"
	"fmt"
)

// func divide(a, b int) (int, error) {

// 	if b == 0 {
// 		return 0, errors.New("can not divide with zero")

// 	}
// 	return a / b, nil

// }

// func main() {

// 	c, d := divide(10, 10)
// 	fmt.Println(c, d)
// }

var outoftea = fmt.Errorf("we are out of tea")
var nopower = fmt.Errorf("can not boil")

func makingtea(a int) error {
	if a == 1 {
		return outoftea
	} else if a == 2 {
		return nopower

	}
	return nil

}

func main() {

	for a := range 4 {

		b := makingtea(a)

		if b != nil {

			if errors.Is(b, outoftea) {
				fmt.Println("need to buy tea")
			} else if errors.Is(b, nopower) {
				fmt.Println("the water is not burning  ")
			}

			fmt.Println("making tea")
		}

	}
}
