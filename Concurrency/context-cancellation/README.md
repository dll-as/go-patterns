# Context Cancellation Pattern (Go Concurrency Pattern)

## What is Context Cancellation?

Context allows:
- Propagating cancellation signals to multiple goroutines
- Setting deadlines / timeouts
- Passing shared values

You use this pattern when:
- You have goroutines doing work for a request
- You want to stop work if the request is canceled
- You want to avoid leaking goroutines

---

## Common Real-World Examples

- HTTP request handling
- Background jobs with timeout
- API calls with deadlines
- Worker pool cancellation

---

## Example in this project

We simulate:
- Multiple workers doing work
- Cancel all workers after a timeout
