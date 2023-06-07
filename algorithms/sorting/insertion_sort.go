package sorting

import "fmt"

func Insertion_Sort(input []int) {
	fmt.Println("Insertion Sort: ")
	var inputLen int = len(input)
	for i := 1; i < inputLen; i++ {
		var curVal, j int = input[i], i - 1
		for j >= 0 && curVal < input[j] {
			input[j+1] = input[j]
			j--
		}
		input[j+1] = curVal
	}
	fmt.Println(input)
}
