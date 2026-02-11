# Table-Driven Tests (Go Idiomatic Pattern)

## What is Table-Driven Testing?

- Define test cases in a table (slice of structs)
- Iterate over them in a single test function
- Avoid repetition and improve readability

Use when:
- Testing a function with multiple scenarios
- Need clear test case organization
- Want to easily add new test cases
