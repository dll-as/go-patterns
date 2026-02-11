package main

import "fmt"

func main() {
	// Slice of empty interfaces
	values := []interface{}{42, "hello", 3.14, true}

	for _, v := range values {
		// Type switch to handle different types
		switch val := v.(type) {
		case int:
			fmt.Println("Integer:", val)
		case string:
			fmt.Println("String:", val)
		case float64:
			fmt.Println("Float:", val)
		case bool:
			fmt.Println("Boolean:", val)
		default:
			fmt.Println("Unknown type")
		}
	}
}
