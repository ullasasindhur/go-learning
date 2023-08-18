package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// deadlock()
	unbuffered()
	buffered()
	goroutine()
}

func deadlock() {
	var data chan int = make(chan int)
	data <- 55
	var n int = <-data
	fmt.Println(n)
}

func unbuffered() {
	fmt.Println("\n\nUnbuffered Function")
	var data chan int = make(chan int)
	fmt.Println("data", data)
	go func() {
		data <- 589
	}()
	var output int = <-data
	fmt.Println("output", output)
}

func buffered() {
	fmt.Println("\n\nBuffered Function")
	var data chan int = make(chan int, 4)
	data <- 1
	data <- 2
	data <- 3
	for i := 0; i < 3; i++ {
		var output int = <-data
		fmt.Println("output -", i, " : ", output)
	}
}

func goroutine() {
	fmt.Println("\n\nGo Routines:")
	var data chan int = make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			data <- i
		}
		close(data)
	}()
	for n := range data {
		fmt.Println(n)
	}
}

func getValue() int {
	return rand.Intn(100)
}
