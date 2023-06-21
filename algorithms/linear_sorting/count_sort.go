package linear_sorting

import "fmt"

func Count_Sort(input []int, inputRange int, inputLen int) {
	fmt.Println("\nInput:")
	fmt.Println(input)
	fmt.Println("\nCount Sort: ")
	var rangeArr []int = make([]int, inputRange)
	for i := 0; i < inputLen; i++ {
		rangeArr[input[i]] += 1
	}
	for i := 1; i < inputRange; i++ {
		rangeArr[i] = rangeArr[i] + rangeArr[i-1]
	}
	var outputArr []int = make([]int, inputLen)
	for i := inputLen - 1; i >= 0; i-- {
		outputArr[rangeArr[input[i]]-1] = input[i]
		rangeArr[input[i]] -= 1
	}
	fmt.Println(outputArr)
}
