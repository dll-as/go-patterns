# Worker Pool Pattern (Go Concurrency Pattern)

## What is Worker Pool?

Worker Pool is a concurrency pattern where a **fixed number of workers** process tasks from a **shared queue**.
It helps **control resource usage** while still achieving concurrency.

You use this pattern when:
- You have **many tasks** to process
- You want to **limit the number of concurrent goroutines**
- Tasks can be processed independently

---

## Common Real-World Examples

- Processing HTTP requests in batches
- Image or video processing
- Job queues / task scheduling
- Crawlers or scrapers

---

## Example in this project

We have a **task queue**:
- Fixed number of workers pick tasks from the queue
- Main waits until all tasks are processed
