# Future / Promise Pattern (Go Concurrency Pattern)

## What is Future / Promise?

Future represents a value that will be available in the future.

- A task runs asynchronously
- It returns immediately with a "future"
- Later, you wait and get the result

In Go, this is commonly implemented using:
- Goroutines
- Channels

You use this pattern when:
- A task takes time (IO, API call, DB query)
- You want non-blocking execution
- You want to get the result later

---

## Common Real-World Examples

- HTTP API calls
- Database queries
- File downloads
- Background computations

---

## Example in this project

We simulate:
- An async API call
- Return a Future (channel)
- Wait for result later

# 🔎 Slightly More Advanced Version (With Error)

In real systems, we usually return **value + error**:

```go
type Result struct {
	Value string
	Err   error
}

func fetchData() <-chan Result {
	ch := make(chan Result)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- Result{Value: "Success", Err: nil}
		close(ch)
	}()

	return ch
}
```