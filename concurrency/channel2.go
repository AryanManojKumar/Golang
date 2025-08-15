package main

import (
	"fmt"
	"time"
)

func main() {

	// go func() {
	// 	for i := range 6 {
	// 		fmt.Println(i)
	// 		ch <- i
	// 		time.Sleep(500 * time.Millisecond)

	// 	}
	// }()

	// go func() {
	// 	b := <-ch
	// 	fmt.Println(b)

	// 	// for b < 6 {
	// 	// 	time.Sleep(500 * time.Millisecond)
	// 	// 	fmt.Println("printing", b)
	// 	// }

	// }()
	// time.Sleep(60 * time.Second)

	// go func() {

	// 	var i int

	// 	for i <= 5 {

	// 		fmt.Println(i)

	// 		go func() {

	// 			r := <-ch
	// 			fmt.Println(r)

	// 		}()
	// 		ch <- i
	// 		i++
	// 	}

	// }()

	// time.Sleep(2 * time.Second)

	ch := make(chan int)
	var f int
	go func() {
		var i int
		for i <= 5 {

			ch <- i

			i++

		}
	}()

	go func() {

		for f < 5 {
			f = <-ch
			fmt.Println(f)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(5 * time.Second)

}
