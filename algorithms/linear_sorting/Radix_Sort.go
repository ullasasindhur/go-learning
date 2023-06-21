package linear_sorting

import "fmt"

func Radix_Sort(input []int, inputLen int) {
	fmt.Println("\nInput:")
	fmt.Println(input)
	fmt.Println("\nRadix Sort:")
	var maxInput int = findMax(input, inputLen)
	for base := 1; maxInput/base > 0; base *= 10 {
		count_sort(input, inputLen, base, 10)
	}
	fmt.Println(input)
}

func findMax(input []int, inputLen int) int {
	var max int = input[0]
	for i := 1; i < inputLen; i++ {
		if input[i] > max {
			max = input[i]
		}
	}
	return max
}

func count_sort(input []int, inputLen int, base int, d int) {
	var rangeArr []int = make([]int, d)
	for i := 0; i < inputLen; i++ {
		rangeArr[(input[i]/base)%10] += 1
	}
	for i := 1; i < d; i++ {
		rangeArr[i] += rangeArr[i-1]
	}
	var tempArr []int = make([]int, inputLen)
	for i := inputLen - 1; i >= 0; i-- {
		var loc int = ((input[i] / base) % 10)
		tempArr[rangeArr[loc]-1] = input[i]
		rangeArr[loc] -= 1
	}
	for i := 0; i < inputLen; i++ {
		input[i] = tempArr[i]
	}
}
