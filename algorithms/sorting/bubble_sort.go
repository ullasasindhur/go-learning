package sorting

import "fmt"

func Bubble_Sort(input []int) {
	fmt.Println("Bubble Sort: ")
	var inputLen int = len(input)
	for i := 0; i < inputLen-1; i++ {
		for j := 0; j < inputLen-1; j++ {
			if input[j] > input[j+1] {
				temp := input[j]
				input[j] = input[j+1]
				input[j+1] = temp
			}
		}
	}
	fmt.Println(input)
}
