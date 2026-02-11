# Memento Pattern (Go)

## What is Memento?

Memento is a behavioral design pattern.
It **captures and externalizes an object's internal state** so it can be restored later, without violating encapsulation.

You use this pattern when:
- You want **undo/redo functionality**
- You need **snapshots of object state**
- You want to **preserve encapsulation**

---

## Common Real-World Examples

- Undo/Redo in text editors
- Game checkpoints / save states
- Version control for objects
- Temporary settings / snapshots

---

## Example in this project

We have a **TextEditor**:
- Can write text
- Save state as Memento
- Restore previous state (undo)
