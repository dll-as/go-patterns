# Empty Interface Pattern (Go Idiomatic Pattern)

## What is Empty Interface?

- `interface{}` can hold **any type**
- Useful for generic containers or flexible functions
- Often combined with type assertion or type switch

Use when:
- You need to store values of different types
- You want flexible function arguments
- Before Go 1.18 generics

---

## Common Real-World Examples

- `map[string]interface{}` for JSON
- Logging systems
- Generic containers
- Flexible function arguments

---

## Example in this project

- Demonstrates storing different types in a slice
- Uses type switch to handle each type
