# Dependency Injection (Constructor-based, Go)

## What is Dependency Injection?

- Pass dependencies into a struct instead of creating them inside
- Constructor function usually handles injection
- Makes testing and mocking easy

You use this pattern when:
- A struct depends on external services
- You want to swap implementations for testing
- You want clean separation of concerns

---

## Common Real-World Examples

- Service using repository or database
- HTTP handler using Logger or Service
- Payment service using PaymentGateway interface
