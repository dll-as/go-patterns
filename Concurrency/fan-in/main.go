package main

import (
	"fmt"
	"time"
)

// Producer function
func producer(id int, ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i + (id-1)*10
		time.Sleep(time.Millisecond * 100) // simulate work
	}

	close(ch)
}

// Fan-In function
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)

	for _, ch := range channels {
		go func(c <-chan int) {
			for val := range c {
				out <- val
			}
		}(ch)
	}

	// Close out channel when all inputs are done
	go func() {
		for _, ch := range channels {
			for range ch {
				// drain channel (already handled in goroutine above)
			}
		}

		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(1, ch1)
	go producer(2, ch2)

	merged := fanIn(ch1, ch2)

	for val := range merged {
		fmt.Println("Received:", val)
	}
}
