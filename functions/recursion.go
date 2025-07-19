package main

import "fmt"

// // 5=5*4*3*2*1
// func fact(a int) int {

// 	if a == 0 {
// 		return 1
// 	}
// 	return a * fact(a-1)
// }

// func main() {
// 	fmt.Println(fact(5))
// }

// 0,1,1,2,3,5,8

func fibo(a int) int {

	if a == 0 {
		return 0

	} else if a == 1 {
		return 1
	}
	return fibo(a-1) + fibo(a-2)

}

func main() {

	// for i := range 10 {

	fmt.Println(fibo(4))

}
