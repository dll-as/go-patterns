# Goroutine per Task (Go Concurrency Pattern)

## What is Goroutine per Task?

This pattern launches **a new goroutine for each independent task**.
It allows tasks to execute concurrently and improves throughput in IO-bound or CPU-bound workloads.

You use this pattern when:
- You have **multiple independent tasks**
- Tasks can run concurrently without blocking each other
- You want **simple concurrency** without complex coordination

---

## Common Real-World Examples

- HTTP request handling in web servers
- Parallel file downloads
- Concurrent computations (e.g., map/reduce)
- Background jobs

---

## Example in this project

We have a list of tasks:
- Each task runs in a separate goroutine
- Main waits for all tasks to complete using WaitGroup
