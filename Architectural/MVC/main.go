package main

import (
	"fmt"
	"net/http"
)

// Model
type User struct {
	ID   int
	Name string
}

// Mock database
var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

// Controller
func userHandler(w http.ResponseWriter, r *http.Request) {
	for _, user := range users {
		// View
		fmt.Fprintf(w, "User: %d, Name: %s\n", user.ID, user.Name)
	}
}

func main() {
	http.HandleFunc("/users", userHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
