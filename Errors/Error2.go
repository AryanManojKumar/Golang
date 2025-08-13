package main

import (
	"fmt"
	"os"
)

func openingfile(a string) error {

	file, err := os.Open(a)

	if err != nil {
		var filedoesntexist = fmt.Errorf("this file doesnt exist")
		return filedoesntexist

	}
	fmt.Println("the file is opening")

	defer file.Close()
	return nil

}

func main() {

	a := openingfile("maps.go")

	fmt.Println(a)

}
