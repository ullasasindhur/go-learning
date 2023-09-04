// In golang we have composition
package inheritance

import "fmt"

type father struct {
	name string
}

func (f father) getFatherDetails() string {
	return f.name
}

type mother struct {
	name string
}

func (m mother) getMotherDetails() string {
	return m.name
}

type child struct {
	father
	mother
}

func (c child) getDetails() string {
	return "mother father"
}

func (c child) getFatherDetails() string {
	c.father.name = "55"
	return c.father.name
}

func Main() {
	fmt.Println("Inheritance:")
	var f = father{name: "father"}
	var m = mother{name: "mother"}
	var child = child{mother: m, father: f}
	fmt.Println(child.getDetails())
	fmt.Println(child.getFatherDetails())
	fmt.Println(child.getMotherDetails())
	fmt.Println(child.father.getFatherDetails())
}
