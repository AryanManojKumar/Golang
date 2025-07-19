package main

import (
	"fmt"
)

func main() {

	const s = "😙😋🥹"

	fmt.Println(s, len(s))

	for idx, runeValue := range s {
		fmt.Printf("ok", runeValue, idx)
	}
	// fmt.Println("rune count", utf8.RuneCountInString(s))

}
