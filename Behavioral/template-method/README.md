# Template Method Pattern (Go)

## What is Template Method?

Template Method is a behavioral design pattern.
It **defines the skeleton of an algorithm in a method**, deferring some steps to subclasses.
This allows subclasses to **override certain steps without changing the overall algorithm structure**.

You use this pattern when:
- You have **algorithms with common structure**
- Some steps **vary between implementations**
- You want **to avoid code duplication**

---

## Common Real-World Examples

- Beverage preparation (Tea, Coffee)
- File processing (CSV, JSON, XML)
- Game AI turns (common steps with variations)
- Report generation (different formats, same workflow)

---

## Example in this project

We have **Beverage preparation**:
- Steps: BoilWater, Brew, Pour, AddCondiments
- Subclasses implement Brew and AddCondiments
