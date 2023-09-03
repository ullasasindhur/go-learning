package main

import "fmt"

func fun(in interface{}) {
	switch in.(type) {
	case int:
		fmt.Println("Type: int", in.(int))
	case string:
		fmt.Println("Type: string", in.(string))
	case float64:
		fmt.Println("Type: float64", in.(float64))
	default:
		fmt.Println("Type not found")
	}
}

func main() {
	fun("Go")
	fun(55)
	fun(50.05)
	fun(true)
}
