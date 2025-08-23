package main

import (
	"fmt"
	"time"
)

func b(id int, work <-chan int) {

	for i := range work {
		fmt.Println("worker", id, "is doing the work", i)
		time.Sleep(2 * time.Second)
		fmt.Println("the work", i, "has been completed by", id)

	}

}

func main() {

	a := make(chan int, 4) //work

	for i := range 4 {
		go b(i, a)
	}

	for j := range 3 {
		a <- j

	}
	close(a)

	time.Sleep(10 * time.Second)

}
