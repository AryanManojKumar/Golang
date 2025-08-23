package main

import (
	"fmt"
	"sync"
	"time"
)

// func ariel(a int) {

// 	fmt.Println(a, "is doing the work")
// 	time.Sleep(2 * time.Second)
// 	fmt.Println(a, "has completed the work")

// }

// func main() {

// 	var a int

// 	for range 4 {
// 		go ariel(a)
// 	}

// 	// time.Sleep(10 * time.Second)
// }

// func okay(a int, wg *sync.WaitGroup) {
// 	fmt.Println(a, "is doing the work")
// 	time.Sleep(2 * time.Second)
// 	fmt.Println(a, "has completed the work")
// 	defer wg.Done()

// }

// func main() {

// 	var wg sync.WaitGroup

// 	for a := range 4 {
// 		wg.Add(1)
// 		go okay(a, &wg)
// 	}

// 	wg.Wait()
// }

func downloader(a int, link <-chan int, wg *sync.WaitGroup) {

	for f := range link {

		fmt.Println(a, "is downloading", f)
		time.Sleep(3 * time.Second)
		fmt.Println(a, "has completed the download of ", f)

	}
	defer wg.Done()

}

func main() {
	var wg sync.WaitGroup
	link := make(chan int, 3)

	for i := range 3 {
		link <- i
	}
	close(link)

	for f := range 2 {
		wg.Add(1)
		go downloader(f, link, &wg)
	}

	wg.Wait()
	fmt.Println("All downloading done")

}
