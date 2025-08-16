package main

import (
	"fmt"
	"time"
)

func worker(a int, ch chan bool) {

	fmt.Println("worker started ", a)
	time.Sleep(time.Second)
	fmt.Println(" worker finished", a)
	ch <- true
}

func main() {

	ch := make(chan bool)

	for i := range 3 {
		go worker(i, ch)

	}

	for i := 0; i < 3; i++ {
		<-ch
	}

}
