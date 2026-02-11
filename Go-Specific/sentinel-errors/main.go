package main

import (
	"errors"
	"fmt"
)

// Define sentinel errors
var (
	ErrUserNotFound     = errors.New("user not found")
	ErrPermissionDenied = errors.New("permission denied")
)

// Mock service
func getUser(id int) error {
	if id == 0 {
		return ErrUserNotFound
	}
	return nil
}

func main() {
	err := getUser(0)
	if err != nil {
		fmt.Println("Error:", err)

		// Check sentinel
		if errors.Is(err, ErrUserNotFound) {
			fmt.Println("Handle user not found")
		}
	}
}
