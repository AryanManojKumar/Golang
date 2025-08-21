package main

import (
	"fmt"
	"time"
)

// func main() {

// a := time.NewTimer(2 * time.Second)

// for range 2 {
// 	select {
// 	case <-a.C:
// 		fmt.Println("time 1 fired")

// 	case <-time.After(3 * time.Second):
// 		fmt.Println("timeout")
// 	}
// }

// 	f := time.NewTimer(2 * time.Second)

// 	go func() {

// 		select {
// 		case <-f.C:
// 			fmt.Println("Timer fired")
// 		}

// 	}()

// 	stop := f.Stop()

// 	if stop == true {
// 		fmt.Println("timer stopped", stop)
// 	}

// 	time.Sleep(5 * time.Second)

// }

func main() {

	a := time.NewTimer(15 * time.Second)

	go func() {
		f, v := <-a.C
		fmt.Println("timer is fired ", f, v)

	}()

	time.Sleep(16 * time.Second)

}
