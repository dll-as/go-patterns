# Chain of Responsibility Pattern (Go)

## What is Chain of Responsibility?

Chain of Responsibility is a behavioral design pattern.
It lets you **pass a request along a chain of handlers**.
Each handler can either **handle it** or **forward it** to the next.

You use this pattern when:
- Multiple objects can handle a request
- You want to **decouple sender and receiver**
- You want flexible handling chains

---

## Common Real-World Examples

- Logging systems (INFO, DEBUG, ERROR)
- Event handling systems
- Validation chains (input validation, security checks)
- Customer support ticket routing

---

## Example in this project

We have a **logging system**:
- Handlers: INFO, DEBUG, ERROR
- Each handler decides if it should log or pass the request
