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

func Main() {
	fmt.Println("Inheritance:")
	var f = father{name: "father"}
	var m = mother{name: "mother"}
	var children = child{mother: m, father: f}
	fmt.Println(children.getDetails())
	fmt.Println(children.getFatherDetails())
	fmt.Println(children.getMotherDetails())
	fmt.Println("Single")
	var single = child{mother: m}
	fmt.Println(single.getFatherDetails())
	fmt.Println(single.getMotherDetails())
	fmt.Println(children.getDetails())
}
