package abstraction

import "fmt"

type Rectangle struct {
	length int
	width  int
}

type Square struct {
	length int
}

type Triangle struct {
	base   int
	height int
}

type Area interface {
	Area() int
}

func (t Triangle) Area() int {
	return t.base * t.height
}

func (s Square) Area() int {
	return s.length * s.length
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func Main() {
	fmt.Println("Abstraction: ")
	var area Area
	var square = Square{length: 5}
	area = square
	fmt.Println(area.Area())
	fmt.Println(square.Area())
	var triangle = Triangle{height: 5, base: 5}
	area = triangle
	fmt.Println(area.Area())
	fmt.Println(triangle.Area())
	fmt.Println(triangle.base)
}
