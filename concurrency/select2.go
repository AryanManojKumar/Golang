package main

import (
	"fmt"
	"time"
)

// func temp(a chan<- string) {

// 	a <- "hot"

// }

// func humidity(a chan<- string) {
// 	time.Sleep(2 * time.Second)
// 	a <- "not humid"
// }

// func motion(a chan<- string) {
// 	time.Sleep(time.Second)

// 	a <- "moving"
// }

// func main() {

// 	a := make(chan string, 3)

// 	go temp(a)
// 	go humidity(a)
// 	go motion(a)

// 	for range 3 {
// 		select {
// 		case ch := <-a:
// 			fmt.Println("this data came ", ch)
// 		case ch2 := <-a:
// 			fmt.Println("this data came ", ch2)
// 		case ch3 := <-a:
// 			fmt.Println("this data came  ", ch3)

// 		}
// 	}

// 	time.Sleep(2 * time.Second)

// }

func main() {

	a := make(chan string)
	b := make(chan string)
	c := make(chan string)

	go func() {

		time.Sleep(time.Second)
		a <- "the temp sensor is sending 100c "

	}()

	go func() {
		time.Sleep(2 * time.Second)
		b <- "humidity sensor says its humid today"

	}()

	go func() {
		time.Sleep(3 * time.Second)
		c <- "motion sensor says something is moving "

	}()

	for range 3 {
		select {
		case ch1 := <-a:
			fmt.Println(ch1)
		case ch2 := <-b:
			fmt.Println(ch2)
		case ch3 := <-c:
			fmt.Println(ch3)
		}
	}

}
