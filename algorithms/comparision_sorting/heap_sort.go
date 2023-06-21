package comparision_sorting

import "fmt"

func Heap_Sort(input []int) {
	fmt.Println("\nHeap Sort: ")
	var inputLen int = len(input)
	for i := (inputLen / 2) - 1; i >= 0; i-- {
		heapify(input, i, inputLen)
	}
	for i := inputLen - 1; i >= 0; i-- {
		var temp int = input[0]
		input[0] = input[i]
		input[i] = temp
		heapify(input[0:i], 0, i)
	}
	fmt.Println(input)
}

func heapify(input []int, index int, inputLen int) {
	var left int = 2*index + 1
	var right int = left + 1
	var min int = index
	if left < inputLen && input[left] > input[min] {
		min = left
	}
	if right < inputLen && input[right] > input[min] {
		min = right
	}
	if min != index {
		var temp int = input[min]
		input[min] = input[index]
		input[index] = temp
		heapify(input, min, inputLen)
	}
}
