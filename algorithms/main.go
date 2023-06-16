package main

import (
	"fmt"
	"math/rand"

	"github.com/ullasasindhur/go-learning/algorithms/sorting"
)

func main() {
	fmt.Println("Welcome to Algorithms Package")
	var input []int = rand.Perm(20)
	fmt.Println("Input: ")
	fmt.Println(input)
	sorting.Insertion_Sort(getCopySlice(input))
	sorting.Bubble_Sort(getCopySlice(input))
	sorting.Selection_Sort(getCopySlice(input))
	sorting.Quick_Sort(getCopySlice(input))
	sorting.Merge_Sort(getCopySlice(input))
}

func getCopySlice(input []int) []int {
	var ouput []int = make([]int, len(input))
	copy(ouput, input)
	return ouput
}
