# Embedding (Composition over Inheritance)

## What is Embedding?

- Use struct embedding instead of inheritance
- Compose behavior by embedding other structs
- Allows "has-a" or "is-like" relationships
- Idiomatic Go approach to reuse code

Use when:
- You want to share fields/methods across types
- Avoid deep inheritance trees
- Prefer composition over inheritance

---

## Common Real-World Examples

- Logger embedded in service structs
- Config struct embedded in multiple components
- HTTP handler embedding common middleware
