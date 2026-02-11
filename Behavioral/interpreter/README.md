# Interpreter Pattern (Go)

## What is Interpreter?

Interpreter is a behavioral design pattern.
It **defines a way to evaluate sentences in a language**.

You use this pattern when:
- You have a **simple language** to interpret
- You need to **evaluate expressions dynamically**
- You want to separate grammar from execution

---

## Common Real-World Examples

- Mathematical expression evaluators
- Query languages or filters
- Rule engines / business rules
- Parsing simple scripts

---

## Example in this project

We have a **simple expression evaluator**:
- Expressions: Numbers and addition
- Each node in the expression tree evaluates itself
