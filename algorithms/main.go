package main

import (
	"fmt"
	"math/rand"

	"github.com/ullasasindhur/go-learning/algorithms/sorting"
)

func main() {
	fmt.Println("Welcome to Algorithms Package")
	var input []int = rand.Perm(20)
	var input2 []int = make([]int, len(input))
	copy(input2, input)
	fmt.Println("Input: ")
	fmt.Println(input)
	sorting.Insertion_Sort(input)
	sorting.Bubble_Sort(input2)
}
