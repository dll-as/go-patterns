## What is Singleton?
Singleton is a creational design pattern.
It makes sure that **only one instance** of a struct exists in the whole application.

This pattern is useful when:
- You need a shared resource (config, logger, database connection)
- You want global access, but controlled creation


## Key Points
- Only one instance is created
- The instance is created lazily (only when needed)
- Thread-safe (important in Go)

## How it works in Go
In Go, Singleton is usually implemented using:
- `sync.Once`
- A private struct
- A public function to get the instance
