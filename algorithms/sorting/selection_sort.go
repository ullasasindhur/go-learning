package sorting

import "fmt"

func Selection_Sort(input []int) {
	fmt.Println("Selection Sort:")
	var inputLen int = len(input)
	for i := 0; i < inputLen-1; i++ {
		var min int = i
		for j := i + 1; j < inputLen; j++ {
			if input[min] > input[j] {
				min = j
			}
		}
		if min != i {
			temp := input[i]
			input[i] = input[min]
			input[min] = temp
		}
	}
	fmt.Println(input)
}
