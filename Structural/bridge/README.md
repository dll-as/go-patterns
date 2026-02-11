# Bridge Pattern (Go)

## What is Bridge?

Bridge is a structural design pattern.
It **decouples an abstraction from its implementation**, so both can vary independently.

You use this pattern when:
- You have multiple abstractions and multiple implementations
- You want to avoid combinatorial explosion of subclasses
- You want flexibility in extending the system

---

## Common Real-World Examples

- Shapes and Colors (Circle-Red, Square-Blue)
- Notification systems (Email/SMS on different platforms)
- GUI toolkits (Button-OSX, Button-Windows)
- Payment systems (Payment type + Gateway)

---

## Example in this project

We have **Notifications**:
- Types: Message, Alert
- Platforms: Email, SMS

Bridge separates **Type** from **Platform**, so any combination is possible.

---

## 🧠 Key Points (Simple)

* `Notification` is **abstraction**
* `Platform` is **implementation**
* Bridge allows **any combination** of abstraction and implementation
* Avoids subclass explosion (instead of `EmailMessage`, `EmailAlert`, `SMSMessage`, `SMSAlert`…)

---

## 🆚 Adapter vs Bridge

| Adapter                           | Bridge                                   |
| --------------------------------- | ---------------------------------------- |
| Make incompatible interfaces work | Decouple abstraction from implementation |
| Wraps an existing object          | Separates interface from implementation  |
| Example: StripeAdapter            | Example: NotificationType + Platform     |