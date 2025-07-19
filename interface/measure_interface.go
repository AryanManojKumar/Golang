type measure interface {
	area() int
	peri() int
}

type rectangle struct {
	width  int
	height int
}

type square struct {
	width int
}

func (r rectangle) area() int {
	return r.width * r.height

}

func (r rectangle) peri() int {
	return 2 * r.width * r.height
}

func (s square) area() int {
	return s.width * s.width
}

func (s square) peri() int {
	return 4 * s.width
}

func measurement(m measure) {
	fmt.Println(m)
	fmt.Println(m.area())
	fmt.Println(m.peri())
}

func main() {
	a := rectangle{width: 10, height: 12}
	b := square{width: 5}

	measurement(a)
	measurement(b)

}