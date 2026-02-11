# Error Type Assertion (Go Idiomatic Pattern)

## What is Error Type Assertion?

- Assert the error to a specific type
- Allows accessing custom fields or methods
- Safer than string matching

Use when:
- Handling errors with structured types
- Need extra info (e.g., path, operation, code)
- Want type-specific handling

---

## Example in this project

- Use custom error type for a file operation
- Demonstrate type assertion
