package sorting

import "fmt"

func Merge_Sort(input []int) {
	fmt.Println("Merge Sort:")
	var inputLen int = len(input)
	merge(input, 0, inputLen-1)
	fmt.Println(input)
}

func merge(input []int, lb int, ub int) {
	if lb < ub {
		var mid int = (lb + ub) / 2
		merge(input, lb, mid)
		merge(input, mid+1, ub)
		merging(input, lb, mid, ub)
	}
}

func merging(input []int, lb int, mid int, ub int) {
	var arr1Len int = mid - lb + 1
	var arr2Len int = ub - mid
	var arr1, arr2 []int = make([]int, arr1Len), make([]int, arr2Len)
	for i := 0; i < arr1Len; i++ {
		arr1[i] = input[lb+i]
	}
	for i := 0; i < arr2Len; i++ {
		arr2[i] = input[mid+i+1]
	}
	var i, j, k int = 0, 0, lb
	for i < arr1Len && j < arr2Len {
		if arr1[i] < arr2[j] {
			input[k] = arr1[i]
			i++
		} else {
			input[k] = arr2[j]
			j++
		}
		k++
	}
	for i < arr1Len {
		input[k] = arr1[i]
		i++
		k++
	}
	for j < arr2Len {
		input[k] = arr2[j]
		j++
		k++
	}
}
