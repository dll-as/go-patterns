# Semaphore Pattern (Go Concurrency Pattern)

## What is Semaphore?

Semaphore is a concurrency control pattern.

It limits:
- How many goroutines can run at the same time.

Mutex = only 1 allowed
Semaphore = N allowed

In Go, semaphore is commonly implemented using:
- Buffered channels

---

## When to Use

- Limit concurrent API calls
- Limit DB connections
- Limit file processing
- Control resource usage

---

## Common Real-World Examples

- Max 10 concurrent HTTP requests
- Max 5 DB queries at once
- Limit parallel downloads
- Limit CPU-heavy tasks

---

## Example in this project

We simulate:
- 10 tasks
- Only 3 can run at the same time
