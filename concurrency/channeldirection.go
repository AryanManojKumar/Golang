package main

import (
	"fmt"
)

func send(ch chan<- int) {

	fmt.Println("this function will only send the data to the pipeline")
	ch <- 6

}

func revicever(ch <-chan int) {
	fmt.Println("this function gets the data")
	a := <-ch
	fmt.Println(a)
}

func main() {

	a := make(chan int)

	go send(a)
	revicever(a)

	// time.Sleep(time.Second)
}
