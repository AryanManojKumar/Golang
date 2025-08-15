package main

import (
	"fmt"
	"time"
)

func main() {

	a := make(chan int)

	go func() {

		fmt.Println("getting value")
		time.Sleep(500 * time.Millisecond)
		a <- 51
		fmt.Println("send")

	}()

	f := <-a
	fmt.Println("got value ", f)

}
