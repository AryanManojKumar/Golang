package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan int)
	b := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)

		a <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)

		b <- 2

	}()

	for range 2 {
		select {
		case ch := <-a:
			fmt.Println("first value recived :", ch)
		case ch2 := <-b:
			fmt.Println("second value recived:", ch2)
		}
	}

}
