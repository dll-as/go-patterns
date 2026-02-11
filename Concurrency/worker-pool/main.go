package main

import (
	"fmt"
	"sync"
	"time"
)

// Task struct
type Task struct {
	ID      int
	Message string
}

// Worker function
func worker(id int, jobs <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for task := range jobs {
		fmt.Printf("Worker %d started task %d: %s\n", id, task.ID, task.Message)
		time.Sleep(time.Millisecond * 500) // simulate work
		fmt.Printf("Worker %d finished task %d\n", id, task.ID)
	}
}

func main() {
	const numWorkers = 3
	const numTasks = 10

	// Create a job queue
	jobs := make(chan Task, numTasks)
	var wg sync.WaitGroup

	// Start fixed number of workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// Send tasks to queue
	for i := 1; i <= numTasks; i++ {
		jobs <- Task{
			ID:      i,
			Message: fmt.Sprintf("Process item %d", i),
		}
	}
	close(jobs) // close channel so workers finish when tasks done

	// Wait for all workers
	wg.Wait()
	fmt.Println("All tasks processed")
}
