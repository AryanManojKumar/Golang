package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int64
	var wg sync.WaitGroup
	var mg sync.Mutex

	for range 10 {
		wg.Add(1)

		go func() {

			for range 1000 {
				mg.Lock()
				counter++
				mg.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("COunter:", counter)

}
