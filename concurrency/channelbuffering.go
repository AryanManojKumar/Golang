package main

import (
	"fmt"
)

func main() {

	a := make(chan string, 3) // buffer size 2

	a <- "buffered"
	a <- "message"
	a <- "third buffer"

	fmt.Println(<-a)
	fmt.Println(<-a)
	fmt.Println(<-a)
}
