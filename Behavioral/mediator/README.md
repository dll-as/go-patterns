# Mediator Pattern (Go)

## What is Mediator?

Mediator is a behavioral design pattern.
It **defines an object that encapsulates how a set of objects interact**, reducing direct dependencies between them.

You use this pattern when:
- Many objects communicate with each other
- You want to **reduce coupling**
- You want **centralized control** of interactions

---

## Common Real-World Examples

- Chat rooms (users communicate via mediator)
- Air traffic control systems
- GUI components interactions
- Event bus or notification systems

---

## Example in this project

We have a **chat room**:
- Users send messages
- Mediator (ChatRoom) delivers messages to all users
- Users do not know about each other directly
