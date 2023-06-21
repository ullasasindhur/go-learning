package main

import (
	"fmt"
	"math/rand"

	comparision_sorting "github.com/ullasasindhur/go-learning/algorithms/comparision_sorting"
	linear_sorting "github.com/ullasasindhur/go-learning/algorithms/linear_sorting"
)

func main() {
	fmt.Print("*****  Welcome to Algorithms Package  *****\n\n")
	var inputLen int = 20
	var input []int = rand.Perm(inputLen)
	fmt.Print("* Comparision Sorting :-\n\n")
	fmt.Println("Input: ")
	fmt.Println(input)
	comparision_sorting.Insertion_Sort(getCopySlice(input))
	comparision_sorting.Bubble_Sort(getCopySlice(input))
	comparision_sorting.Selection_Sort(getCopySlice(input))
	comparision_sorting.Quick_Sort(getCopySlice(input))
	comparision_sorting.Merge_Sort(getCopySlice(input))
	comparision_sorting.Heap_Sort(getCopySlice(input))
	fmt.Println("\n\n* Linear Sorting :-")
	var inputRange int = 7
	for i := 0; i < inputLen; i++ {
		input[i] = rand.Intn(inputRange)
	}
	fmt.Println("\nInput:")
	fmt.Println(input)
	linear_sorting.Count_Sort(getCopySlice(input), inputRange, inputLen)
}

func getCopySlice(input []int) []int {
	var ouput []int = make([]int, len(input))
	copy(ouput, input)
	return ouput
}
