# Flyweight Pattern (Go)

## What is Flyweight?

Flyweight is a structural design pattern.
It **reduces memory usage** by sharing common data between similar objects.

You use this pattern when:
- You have **many similar objects**
- Most of their state is **intrinsic (shared)**
- Only a small part is **extrinsic (unique per object)**

---

## Common Real-World Examples

- Text editor characters (font, size shared, position unique)
- Game trees (enemy types, bullets, particles)
- Icons or images in UI
- Connection pooling (shared objects)

---

## Example in this project

We have a **character rendering system**:
- Most character info (font, style) is shared
- Position is unique
- Flyweight reduces memory for many characters
