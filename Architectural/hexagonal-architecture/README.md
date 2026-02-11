# Hexagonal Architecture (Ports & Adapters) in Go

## What is Hexagonal Architecture?

- Separate core domain logic from external dependencies
- **Ports**: interfaces defining required operations
- **Adapters**: concrete implementations interacting with external systems
- Core domain knows only about Ports

Use when:
- Building flexible and testable applications
- Want to swap infrastructure easily
- Reduce coupling between domain and frameworks
