package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	// Acquire semaphore
	sem <- struct{}{}
	fmt.Printf("Worker %d started\n", id)

	time.Sleep(time.Second) // simulate work

	fmt.Printf("Worker %d finished\n", id)

	// Release semaphore
	<-sem
}

func main() {
	const maxConcurrent = 3
	const totalTasks = 10

	// Buffered channel = semaphore
	sem := make(chan struct{}, maxConcurrent)

	var wg sync.WaitGroup

	for i := 1; i <= totalTasks; i++ {
		wg.Add(1)
		go worker(i, sem, &wg)
	}

	wg.Wait()
	fmt.Println("All tasks done")
}
