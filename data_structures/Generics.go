package main

import "fmt"

// func value(a any) {

// 	fmt.Println(a)

// }

// func main() {
// 	value("hatunamatata")
// }

type list[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next  *element[T]
	value T
}

func (a *list[T]) push(v T) {
	if a.tail == nil {
		a.head = &element[T]{value: v}
		a.tail = a.head

	} else {
		a.tail.next = &element[T]{value: v}
		a.tail = a.tail.next
	}
}

func (a *list[T]) allements() []T {

	var b []T

	for e := a.head; e != nil; e = e.next {

		b = append(b, e.value)

	}
	return b
}

func main() {

	var f list[int]

	f.push(10)
	f.push(20)

	fmt.Println(f.allements())
}
