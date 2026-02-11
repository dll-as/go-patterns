# Fan-Out Pattern (Go Concurrency Pattern)

## What is Fan-Out?

Fan-Out is a concurrency pattern where **multiple goroutines read from the same channel** to process tasks concurrently.
It allows **parallel processing** of tasks using multiple workers.

You use this pattern when:
- You have **tasks that can be processed independently**
- You want **to increase throughput**
- You want **to balance load among multiple goroutines**

---

## Common Real-World Examples

- Worker Pool pattern (tasks processed by multiple workers)
- Parallel HTTP requests
- Image/video processing in parallel
- Data pipelines

---

## Example in this project

We have **one job channel**:
- Multiple workers read from the same channel
- Each worker processes jobs concurrently
