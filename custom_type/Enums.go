package main

import (
	"fmt"
)

// type of the day  Enums
type day int

const (
	Monday day = iota
	tuesday
	wednesday
	thurday
	friday
	saturday
	sunday
)

func main() {
	today := sunday

	fmt.Println(today)
	fmt.Println(Monday)

	switch today {

	case Monday:
		fmt.Println("monday")
	case tuesday:
		fmt.Println("tuesday")
	case wednesday:
		fmt.Println("wednesday")
	case thurday:
		fmt.Println("thursday")
	case friday:
		fmt.Println("friday")
	case saturday:
		fmt.Println("saturday")
	case sunday:
		fmt.Println("holiday lets goo")

	}

}
