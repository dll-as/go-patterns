# Mutex Pattern (Go Concurrency Pattern)

## What is Mutex?

Mutex (Mutual Exclusion) is a synchronization primitive.

It ensures that:
- Only one goroutine can access a shared resource at a time.

Without Mutex:
- You may get race conditions
- Data may become corrupted

In Go, Mutex is provided by:
- sync.Mutex

---

## When to Use

- Shared counters
- Shared maps
- In-memory caches
- Shared state in services

---

## Common Real-World Examples

- Incrementing request counters
- Updating shared configuration
- Writing to shared memory structure
- Protecting critical sections

---

## Example in this project

We simulate:
- Multiple goroutines incrementing a shared counter
- Mutex ensures safe access
