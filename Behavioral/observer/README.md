# Observer Pattern (Go)

## What is Observer?

Observer is a behavioral design pattern.
It defines a **one-to-many dependency** between objects so that **when one object changes, all its dependents are notified automatically**.

You use this pattern when:
- You have a **subject whose state changes**
- You want multiple observers to react
- You want **loose coupling** between subject and observers

---

## Common Real-World Examples

- GUI frameworks (button click, text change)
- Event listeners / Event buses
- Stock price updates / notification systems
- Publish-subscribe systems

---

## Example in this project

We have a **NewsPublisher**:
- Subscribers are notified when new news arrives
- Subscribers implement `Update()` method
