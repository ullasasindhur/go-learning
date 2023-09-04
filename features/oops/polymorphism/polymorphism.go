package polymorphism

import "fmt"

type Poly interface {
	Length() int
}

type name struct {
	cm int
}

func (n name) Length() int {
	return n.cm
}

type time struct {
	m int
}

func (t time) Length() int {
	return t.m
}

func Main() {
	fmt.Println("Polymorphism:")
	var t = time{m: 1}
	var n = name{cm: 2}
	var face Poly
	face = t
	fmt.Println(face.Length())
	face = n
	fmt.Println(face.Length())
}
