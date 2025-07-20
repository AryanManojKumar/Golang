package main

import (
	"fmt"
)

type severstate int

const (
	stateidle severstate = iota
	stateconnected
	stateerror
	statereturing
)

var state = map[severstate]string{
	stateidle:      "idle",
	stateconnected: "connected",
	stateerror:     "error",
	statereturing:  "retrying",
}

func (s severstate) b() string {
	return state[s]

}

func transition(c severstate) string {
	switch c {
	case stateidle:
		return stateconnected.b()
	case stateconnected, statereturing:
		return stateidle.b()
	case stateerror:
		return stateerror.b()
	default:
		panic(fmt.Errorf("unexpected error"))
	}
}

func main() {
	a := stateconnected
	fmt.Println(a.b())

	n1 := transition(stateconnected)
	fmt.Println("next state will be ", n1)

}
