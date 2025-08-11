package main

import "fmt"

type element struct {
	value int
	next  *element
}
type list struct {
	head *element
	tail *element
}

func (a *list) pushing(v int) {
	if a.tail == nil {
		a.head = &element{value: v}
		a.tail = a.head
	} else {
		a.tail.next = &element{value: v}
		a.tail = a.tail.next
	}
}

func (b *list) fulllist() []int {
	k := []int{}

	for e := b.head; e != nil; e = e.next {

		k = append(k, e.value)

	}

	return k
}

func main() {

	f := list{}

	f.pushing(10)
	f.pushing(20)
	f.pushing(30)
	fmt.Println(f.fulllist())

}
