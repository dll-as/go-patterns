# Graceful Shutdown (Go Idiomatic Pattern)

## What is Graceful Shutdown?

- Stop accepting new requests
- Wait for ongoing tasks to complete
- Release resources (connections, files, goroutines)

Use when:
- Running servers or background workers
- Want clean shutdown on SIGINT/SIGTERM
- Prevent data loss or resource leaks
