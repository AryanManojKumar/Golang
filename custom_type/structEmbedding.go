package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type employee struct {
	person
	Employee_id int
}

func main() {

	a := employee{
		person:      person{name: "abc", age: 22},
		Employee_id: 1,
	}

	var b int
	fmt.Print("Type the employee ID to search for user and age: ")
	fmt.Scanf("%d", &b)

	if b == a.Employee_id {
		fmt.Println("the employee name is  and age is  ", a.name, a.age)
	} else {
		fmt.Println("not in the records")
	}
}
