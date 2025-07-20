package main

import (
	"fmt"
)

// type person struct {
// 	name string
// 	age  int
// }

// func main() {

// 	p1 := person{name: "tushar", age: 69}
// 	p2 := person{name: "aryan", age: 69}

// 	fmt.Println("first persons name is ", p1.name)
// 	fmt.Println("second person's age ", p2.name)

// }

// type person struct {
// 	name string
// 	age  int
// }

// func (a person) laugh() {
// 	fmt.Println(a.name, "is hahaha")

// }

// func main() {
// 	b := person{name: "tushar", age: "69"}

// 	fmt.Println("ur mom so fat")
// 	b.laugh()
// }

// func newperson(nam string) *person {
// 	p := person{name: nam}
// 	p.age = 49
// 	return &p

// }

// func main() {

// 	fmt.Println(person{})

// 	fmt.Println(&person{name: "ary", age: 4})
// 	fmt.Println(person{})

// }

// type books struct {
// 	title  string
// 	author string
// 	pages  int
// 	avl    bool
// }

// func (b *books) printBookDetails() {
// 	fmt.Println(b.title)
// 	fmt.Println(b.author)
// 	fmt.Println(b.pages)
// 	fmt.Println(b.avl)

// }

// func main() {

// 	a := books{
// 		title:  "abc",
// 		author: "tushar",
// 		pages:  69,
// 		avl:    false}

// 	a.printBookDetails()

// }

type student struct {
	rollno int
	name   string
	grades float64
	pass   bool
}

func (a *student) gradeupdate() {
	fmt.Println("student grade is ", a.grades)
	if a.grades >= 50 {
		a.pass = true

	} else {
		a.pass = false
	}
}

func main() {

	b := student{
		rollno: 5,
		name:   "gg",
		grades: 51,
	}
	b.gradeupdate()

	fmt.Println(b.pass)
}
