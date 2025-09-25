package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("Test.txt")
	if err != nil {
		log.Fatal("couldnt open the file", err)
	}
	fmt.Println(data) //this will only print the []byte data of the file

	fmt.Println(string(data)) //this will convert the data into string will only work if the file contains text data

	file, err := os.Open("Test.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	filestats, err := file.Stat() //this'll give the metadata about the file to filestats
	if err != nil {
		panic(err)
	}

	fmt.Println("file size ", filestats.Size())

	// you can even map  all the information that will get and filter out what all you need

	details := map[string]interface{}{
		"name":    filestats.Name(),
		"modtime": filestats.ModTime(),
		"mode":    filestats.Mode(),
	}

	for key, value := range details {
		fmt.Println(key, value)
	}
}
