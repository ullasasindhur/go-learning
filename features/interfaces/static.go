package main

import "fmt"

type static interface {
	area() int
	perimeter() int
}

type params struct {
	len int
	wid int
}

func (p params) area() int {
	return p.len + p.wid
}

func (p params) perimeter() int {
	return 2 * (p.len + p.wid)
}

type parameters struct {
	length  int
	breadth int
}

func (p parameters) area() int {
	return p.length * p.breadth
}

func (p parameters) perimeter() int {
	return 2 * p.breadth * p.length
}

func main() {
	var s static
	s = parameters{length: 5, breadth: 5}
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
	s = params{len: 10, wid: 10}
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}
