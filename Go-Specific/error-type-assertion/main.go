package main

import (
	"fmt"
	"os"
)

// Custom error type
type MyFileError struct {
	Op   string
	Path string
	Err  error
}

func (e *MyFileError) Error() string {
	return fmt.Sprintf("%s %s: %v", e.Op, e.Path, e.Err)
}

// Function that returns a typed error
func readFile(filename string) error {
	if filename == "" {
		return &MyFileError{
			Op:   "read",
			Path: filename,
			Err:  os.ErrNotExist,
		}
	}
	return nil
}

func main() {
	if err := readFile(""); err != nil {
		fmt.Println("Error:", err)

		// Type assertion
		if fe, ok := err.(*MyFileError); ok {
			fmt.Printf("Caught MyFileError: Op=%s, Path=%s, Original=%v\n", fe.Op, fe.Path, fe.Err)
		}
	}
}
