# Functional Options Pattern (Go Idiomatic)

## What is Functional Options?

- Allows building flexible constructors
- Pass configuration as functions
- Avoid many constructor variations

You use this pattern when:
- Your struct has optional fields
- You want readable initialization
- You want extensibility without breaking API

---

## Common Real-World Examples

- http.Client initialization
- Logger configuration
- Database connections
- Any configurable service

---

## Example in this project

We simulate:
- A Server struct
- Options for Port, Timeout, and Debug
