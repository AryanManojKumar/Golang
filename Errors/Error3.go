package main

import (
	"fmt"
)

type argrr struct {
	arg int
	err string
}

// type error interface{
// 	errortext()

// }

func (a *argrr) Error() string {
	return fmt.Sprintf("%d - %s", a.arg, a.err)
}

func f(a *argrr) (int, error) {
	b := string("this wont work")
	if a.arg == 1 {

		return a.arg, &argrr{arg: a.arg, err: b}
	}

	return -1, nil
}

func main() {
	a := argrr{arg: 5, err: "nil"}
	_, b := f(&a)

	if b != nil {
		fmt.Println(b)

	} else {
		fmt.Println("no error")
	}
}
