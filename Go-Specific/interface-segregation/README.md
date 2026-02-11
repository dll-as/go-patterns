# Interface Segregation (Go-style)

## What is Interface Segregation?

- Keep interfaces small and focused
- Avoid forcing implementations to include unnecessary methods
- Each client depends only on the methods it uses

Go style:
- Small, focused interfaces
- Often a single method per interface
- Compose interfaces as needed

---

## Common Real-World Examples

- io.Reader, io.Writer, io.Closer
- Database operations: Queryer, Inserter, Updater
- Logger: InfoLogger, DebugLogger, ErrorLogger

---

## Example in this project

We simulate:
- Different devices
- Each implements only the interfaces they need
- Small, focused interfaces
