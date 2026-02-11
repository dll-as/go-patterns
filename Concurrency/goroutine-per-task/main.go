package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Task %d started\n", id)
	time.Sleep(time.Second) // simulate work
	fmt.Printf("Task %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	tasks := 5

	for i := 1; i <= tasks; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All tasks completed")
}
