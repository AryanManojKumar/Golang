package main

import (
	"fmt"
)

type stats struct {
	health int
	mana   int
}

func (a stats) PrintStats() {
	fmt.Println(a.health, a.mana)

}

type character struct {
	name      string
	classtype string
	stats
}

func (b character) Introduce() {
	fmt.Println(b)
}

func main() {
	aryan := character{
		name: "aryan", classtype: "Coder",
		stats: stats{health: 100, mana: 500}}

	aryan.Introduce()
}
