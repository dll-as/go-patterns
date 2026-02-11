# Builder Pattern (Go)

## What is Builder?

Builder is a creational design pattern.
It is used to **build complex objects step by step**.

You use this pattern when:
- An object has many optional fields
- Constructors become too large or confusing
- You want readable and flexible object creation

---

## Common Real-World Examples

- HTTP servers (port, timeout, TLS, middleware)
- Database configurations
- Cloud service clients
- Request builders (HTTP, SQL)

---

## Example in this project

We build an **HTTP Server config** with optional settings:
- Port
- Timeout
- TLS
- Logging

## 🧠 Key Points (Simple)

* Builder avoids big constructors
* Method chaining makes code readable
* Default values are easy to manage
* Very common in Go projects

---

## 🧪 Go Idiomatic Note

In real Go projects, Builder is often implemented as:

* **Functional Options Pattern**
  (we will cover it later — it is basically a Go-style Builder)

---

## ❌ Without Builder (Problem)

```go
NewServer(8080, 30, false, true, true, ...)
```

Hard to read and easy to break.

---

## ✅ With Builder (Benefit)
```go
NewServerBuilder().
	EnableTLS().
	SetPort(9090).
	Build()
```

Clear and safe.
