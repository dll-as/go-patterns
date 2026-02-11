# Prototype Pattern (Go)

## What is Prototype?

Prototype is a creational design pattern.
It allows you to **copy existing objects** instead of creating new ones from scratch.

You use this pattern when:
- Object creation is expensive
- Objects are similar with small differences
- You want to avoid complex initialization logic

---

## Common Real-World Examples

- Cloning configuration objects
- HTTP request templates
- Game characters or items
- Cache-based object creation

---

## Example in this project

We have a **base server config**.
We clone it and change only what we need.

## 🧠 Key Points (Simple)

* `Clone()` creates a copy of an object
* No need to rebuild from scratch
* Base object stays unchanged
* Shallow copy is often enough in Go

---

## ⚠️ Important Go Note

* `copy := *s` creates a **shallow copy**
* If struct has slices, maps, or pointers → you need **deep copy**

Example:

```go
newSlice := append([]int{}, oldSlice...)
```

---

## 🆚 Builder vs Prototype

| Builder                 | Prototype                 |
| ----------------------- | ------------------------- |
| Builds step by step     | Copies an existing object |
| Good for many options   | Good for similar objects  |
| New instance every time | Clone and modify          |
