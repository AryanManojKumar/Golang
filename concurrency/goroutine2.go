package main

import (
	"fmt"
	"time"
)

func print(which string) {
	for i := range len(which) {
		fmt.Println(which, ":", i)
	}
	time.Sleep(500 * time.Millisecond)
}

func main() {
	go print("holaamigo")
	go print("espanional")

	go func(a string) {
		fmt.Println(a)
	}("realmadrid")

	time.Sleep(2 * time.Second) //so that we can have time to make the routine run
}
