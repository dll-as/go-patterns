// sync.Once guarantees the code runs only once, even with multiple goroutines
// instance is created only on the first call
// This is the most idiomatic Singleton implementation in Go

package main

import (
	"fmt"
	"sync"
)

// singleton struct
type singleton struct {
	value string
}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance returns the single instance of singleton
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{
			value: "I am the only instance",
		}
	})

	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Println(s1.value)
	fmt.Println(s2.value)

	// both variables point to the same instance
	fmt.Println(s1 == s2) // true
}
