# Sentinel Errors (Go Idiomatic Pattern)

## What is Sentinel Error?

- Predefined error variables
- Can be checked anywhere in the code
- Usually exported

Use when:
- You want consistent error handling
- Clients can check specific error
- Simple error signaling is enough

---

## Example in this project

- Define sentinel errors for a simple user service
