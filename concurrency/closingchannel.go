package main

import (
	"fmt"
)

// func main() {

// 	ch := make(chan string, 3)

// 	ch <- "hello"
// 	ch <- "its "
// 	ch <- "me"

// 	for i := range ch {
// 		fmt.Println(i)
// 	}

// }

// func main() {

// 	ch := make(chan int, 1)

// 	ch <- 4
// 	close(ch)

// 	a, ok := <-ch

// 	fmt.Println(a, ok)
// }

// func main() {

// 	a := make(chan int, 5)

// 	go func() {
// 		for i := 0; i < 5; i++ {

// 			a <- i
// 			fmt.Println("sending values", i)

// 		}
// 		close(a)
// 	}()

// 	for v := range a {

// 		fmt.Println("receiving value", v)
// 	}

// }

func main() {

	jobs := make(chan int, 4)
	done := make(chan bool)

	go func() {

		for {
			a, work := <-jobs

			if work {
				fmt.Println("jobs recived", a)

			} else {
				fmt.Println("no jobs available")
				done <- true
				return
			}

		}
	}()

	for i := 0; i <= 3; i++ {
		jobs <- i
		fmt.Println("jobs send", i)

	}
	close(jobs)
	<-done

	_, k := <-jobs
	fmt.Println("recived more jobs", k)

}
