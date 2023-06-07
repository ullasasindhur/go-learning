package main

import (
	"fmt"
	"math/rand"

	"github.com/ullasasindhur/go-learning/algorithms/sorting"
)

func main() {
	fmt.Println("Welcome to Algorithms Package")
	var input []int = rand.Perm(20)
	fmt.Println("Input:")
	fmt.Println(input)
	sorting.Insertion_Sort(input)
}
