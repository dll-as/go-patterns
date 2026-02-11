package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, barrier *sync.WaitGroup) {
	fmt.Printf("Worker %d doing preparation\n", id)
	time.Sleep(time.Millisecond * time.Duration(200*id))

	fmt.Printf("Worker %d reached barrier\n", id)

	// Wait at barrier
	barrier.Done()
	barrier.Wait()

	fmt.Printf("Worker %d passed barrier\n", id)
}

func main() {
	const numWorkers = 5

	var barrier sync.WaitGroup
	barrier.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, &barrier)
	}

	// Wait enough time so program doesn't exit early
	time.Sleep(3 * time.Second)
}
