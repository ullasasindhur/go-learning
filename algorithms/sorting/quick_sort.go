package sorting

import "fmt"

func Quick_Sort(input []int) {
	fmt.Println("Quick Sort: ")
	quick(input, 0, len(input)-1)
	fmt.Println(input)
}

func quick(input []int, lb int, ub int) {
	if lb < ub {
		var location int = partition(input, lb, ub)
		quick(input, lb, location-1)
		quick(input, location+1, ub)
	}
}

func partition(input []int, lb int, ub int) int {
	var pivot int = input[lb]
	var start, end int = lb, ub
	for start < end {
		for start <= ub && pivot >= input[start] {
			start++
		}
		for end >= lb && pivot < input[end] {
			end--
		}
		if start < end {
			var temp int = input[start]
			input[start] = input[end]
			input[end] = temp
		}
	}
	var temp int = input[end]
	input[end] = input[lb]
	input[lb] = temp
	return end
}
