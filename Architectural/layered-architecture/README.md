# Layered Architecture in Go

## What is Layered Architecture?

- Organize code into multiple layers with clear responsibilities
- Each layer communicates with the one below it
- Typical layers:
  1. **Presentation:** HTTP handlers, CLI, UI
  2. **Application:** orchestrates use cases
  3. **Domain:** business models and rules
  4. **Infrastructure:** DB, external services

Use when:
- Building medium-to-large applications
- Want separation of concerns
- Need maintainable and testable structure
