package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "sensor sending data"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)

	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}

}
