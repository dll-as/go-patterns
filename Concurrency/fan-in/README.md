# Fan-In Pattern (Go Concurrency Pattern)

## What is Fan-In?

Fan-In is a concurrency pattern where **multiple channels are merged into a single channel**.
This allows a single receiver to process messages from multiple sources.

You use this pattern when:
- You have **multiple producers**
- You want to **collect results in a single place**
- You want **simpler consumption logic**

---

## Common Real-World Examples

- Multiple workers producing results
- Collecting logs from multiple goroutines
- Aggregating API responses
- Event streams merging

---

## Example in this project

We have **two producers**:
- Each sends integers to its own channel
- Fan-In merges them into a single channel
