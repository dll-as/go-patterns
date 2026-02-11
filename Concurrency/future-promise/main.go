package main

import (
	"fmt"
	"time"
)

// Future function
func fetchData() <-chan string {
	result := make(chan string)

	go func() {
		time.Sleep(2 * time.Second) // simulate long task
		result <- "Data received from API"
		close(result)
	}()

	return result
}

func main() {
	fmt.Println("Requesting data...")

	future := fetchData()

	fmt.Println("Doing other work...")

	// Later, wait for result
	data := <-future

	fmt.Println("Result:", data)
}
