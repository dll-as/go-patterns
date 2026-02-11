package main

import "fmt"

// Stage 1: Generate numbers
func generate(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

// Stage 2: Square numbers
func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

// Stage 3: Print results
func main() {
	// Build pipeline
	numbers := generate(1, 2, 3, 4, 5)
	squared := square(numbers)

	// Consume final output
	for result := range squared {
		fmt.Println(result)
	}
}
