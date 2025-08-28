package main

import (
	"fmt"
	"os"
)

func creatingfile(p string) *os.File {
	fmt.Println("Creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writingfile(f *os.File) {
	fmt.Println("writing file")
	fmt.Fprint(f, "whatever data u want bro")

}

func closingfile(f *os.File) {
	fmt.Println("closing file")
	err := f.Close()
	if err != nil {
		panic(err)
	}

}

func main() {

	p := creatingfile("C:/Users/Aryan/Desktop/Golang/gowithexample/newfile")

	writingfile(p)
	defer closingfile(p)

}
