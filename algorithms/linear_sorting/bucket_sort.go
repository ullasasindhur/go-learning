package linear_sorting

import "fmt"

func Bucket_Sort(input []float64, inputLen int) {
	fmt.Println("\nInput:")
	fmt.Println(input)
	fmt.Println("\nBucket Sort:")
	var d int = 10
	var tempArray [][]float64 = make([][]float64, d)
	for i := 0; i < d; i++ {
		tempArray[i] = make([]float64, 0)
	}
	for i := 0; i < inputLen; i++ {
		var index int = int(input[i] * float64(d))
		tempArray[index] = append(tempArray[index], input[i])
	}
	for i := 0; i < d; i++ {
		if len(tempArray[i]) > 0 {
			insertion_sort(tempArray[i])
		}
	}
	var index int = 0
	for i := 0; i < d; i++ {
		for j := 0; j < len(tempArray[i]); j++ {
			input[index] = tempArray[i][j]
			index++
		}
	}
	fmt.Println(input)
}

func insertion_sort(input []float64) {
	var length int = len(input)
	for i := 1; i < length; i++ {
		curVal, j := input[i], i-1
		for j >= 0 && curVal < input[j] {
			input[j+1] = input[j]
			j--
		}
		input[j+1] = curVal
	}
}
