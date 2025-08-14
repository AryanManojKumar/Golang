package main

import (
	"fmt"
	"time"
)

func printingnumber() {
	for i := range 5 {
		fmt.Println("number", i)
		time.Sleep(500 * time.Millisecond)
	}
	// time.Sleep(500*time.Microsecond)
}

func printingletter() {
	for ch := 'A'; ch <= 'E'; ch++ {
		fmt.Println("letters", ch)
		time.Sleep(500 * time.Millisecond)
	}

}

func main() {

	go printingnumber()
	go printingletter()

	time.Sleep(3 * time.Second)

}
