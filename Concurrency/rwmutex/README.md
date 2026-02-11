# Read-Write Lock (RWMutex) Pattern (Go Concurrency)

## What is RWMutex?

RWMutex = Read-Write Mutex

- Multiple goroutines can **read** at the same time
- Only one goroutine can **write** at a time
- Write blocks readers; readers block writer

Use when:
- Shared resource is **read often but rarely written**
- Optimize concurrency for reads

---

## Common Real-World Examples

- In-memory cache
- Configuration object
- Shared map of counters
- Frequently read data structure

---

## Example in this project

We simulate:
- Multiple readers reading a shared map
- Writers updating the map
- RWMutex ensures safety and better read concurrency
