package main

import "fmt"

func fun(a interface{}) {
	val := a.(string)
	fmt.Println("Value: ", val)
}
func main() {
	var val interface{} = "UllasaSindhur"
	fun(val)
}
