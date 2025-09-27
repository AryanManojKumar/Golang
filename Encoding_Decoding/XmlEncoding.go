package main

import (
	"encoding/xml"
	"log"
	"os"
)

type person struct {
	XMLname xml.Name `xmlname:"person"`
	Name    string   `xml:"name"`
	Uid     int      `xml:"userid"`
}

func main() {

	p := person{
		Name: "APdillion",
		Uid:  23,
	}

	data, err := xml.MarshalIndent(p, " ", " ")
	if err != nil {
		panic(err)
	}

	file, err := os.Create("Person.xml")
	if err != nil {
		panic(err)
	}

	_, err = file.Write(data)
	if err != nil {
		log.Println("could write in file", err)
	}

}
