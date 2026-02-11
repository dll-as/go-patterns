# Error Wrapping (Go Idiomatic Pattern)

## What is Error Wrapping?

- Add context to an error while preserving original error
- Allows checking the original error later
- Improves debugging and logging

Use when:
- Propagating errors up the call stack
- Need additional context
- Want to distinguish error types

---

## Common Real-World Examples

- Database errors
- HTTP request failures
- File operations
- Service layer error handling
