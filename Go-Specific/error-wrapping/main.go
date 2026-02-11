package main

import (
	"errors"
	"fmt"
)

// Original error
var ErrFileNotFound = errors.New("file not found")

func readFile(filename string) error {
	// Simulate error
	if filename == "" {
		return fmt.Errorf("readFile failed: %w", ErrFileNotFound)
	}
	return nil
}

func processFile(filename string) error {
	if err := readFile(filename); err != nil {
		return fmt.Errorf("processFile: cannot process %s: %w", filename, err)
	}
	return nil
}

func main() {
	err := processFile("")
	if err != nil {
		fmt.Println("Error:", err)

		// Check original error
		if errors.Is(err, ErrFileNotFound) {
			fmt.Println("Original error was ErrFileNotFound")
		}
	}
}
