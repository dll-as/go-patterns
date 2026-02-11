package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Millisecond * 200) // simulate work
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start multiple workers (Fan-Out)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All jobs processed")
}
