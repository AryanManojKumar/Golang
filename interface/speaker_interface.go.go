type speaker interface {
	speak()
}

type dog struct{}
type human struct{}

func (d dog) speak() {
	fmt.Println("woof")
}

func (h human) speak() {
	fmt.Println("hello")
}

func saysomething(s speaker) {
	s.speak()

}

func main() {
	a := dog{}

	saysomething(a)
}
