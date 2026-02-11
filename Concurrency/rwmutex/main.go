package main

import (
	"fmt"
	"sync"
)

type SharedData struct {
	mu   sync.RWMutex
	data map[string]int
}

func (s *SharedData) Read(key string) int {
	s.mu.RLock()         // acquire read lock
	defer s.mu.RUnlock() // release read lock
	return s.data[key]
}

func (s *SharedData) Write(key string, value int) {
	s.mu.Lock()         // acquire write lock
	defer s.mu.Unlock() // release write lock
	s.data[key] = value
}

func main() {
	shared := &SharedData{
		data: make(map[string]int),
	}

	var wg sync.WaitGroup

	// Writers
	for i := range 3 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			shared.Write(fmt.Sprintf("key%d", i), i*10)
			fmt.Printf("Writer %d wrote key%d\n", i, i)
		}(i)
	}

	// Readers
	for i := range 5 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i%3)
			val := shared.Read(key)
			fmt.Printf("Reader %d read %s = %d\n", i, key, val)
		}(i)
	}

	wg.Wait()
}
