package main

import (
	"fmt"
)

func main() {

	type request struct {
		add int
		get chan int
	}

	reqchan := make(chan request)

	go func() {
		count := 0

		for req := range reqchan {
			if req.add != 0 {
				count += req.add

			}
			if req.get != nil {
				req.get <- count

			}

		}

	}()

	reqchan <- request{add: 5}
	getchan := make(chan int)
	// reqchan <- request{get: getchan}
	a := request{add: 5, get: getchan}
	reqchan <- a

	fmt.Println(<-getchan)

}
