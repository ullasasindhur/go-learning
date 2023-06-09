package main

import (
	"fmt"
	"math/rand"

	"github.com/ullasasindhur/go-learning/algorithms/sorting"
)

func main() {
	fmt.Println("Welcome to Algorithms Package")
	var input []int = rand.Perm(20)
	var inputLen int = len(input)
	var input2 []int = make([]int, inputLen)
	var input3 []int = make([]int, inputLen)
	copy(input2, input)
	copy(input3, input)
	fmt.Println("Input: ")
	fmt.Println(input)
	sorting.Insertion_Sort(input)
	sorting.Bubble_Sort(input2)
	sorting.Selection_Sort(input3)
}
