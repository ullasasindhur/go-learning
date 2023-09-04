package encapsulation

import "fmt"

var num int = 50

func getNum() int {
	return num
}

func setNum(n int) int {
	num = n
	return num
}

func Main() {
	fmt.Println("Encapsulation: ")
	fmt.Println(num)
	fmt.Println(getNum())
	fmt.Println(setNum(55))
}
