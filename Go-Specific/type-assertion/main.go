package main

import "fmt"

func main() {
	var val interface{} = 42

	// Type assertion
	if intValue, ok := val.(int); ok {
		fmt.Println("Integer value:", intValue)
	} else {
		fmt.Println("Not an int")
	}

	// Another example
	val = "hello"
	strValue := val.(string)
	fmt.Println("String value:", strValue)

	// Using type switch
	values := []interface{}{1, "two", true}
	for _, v := range values {
		switch v := v.(type) {
		case int:
			fmt.Println("Int:", v)
		case string:
			fmt.Println("String:", v)
		case bool:
			fmt.Println("Bool:", v)
		default:
			fmt.Println("Unknown type")
		}
	}
}
