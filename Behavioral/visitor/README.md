# Visitor Pattern (Go)

## What is Visitor?

Visitor is a behavioral design pattern.
It **lets you separate an algorithm from the objects on which it operates**.
You can **add new operations without changing the object structure**.

You use this pattern when:
- You have **different types of elements in a structure**
- You want to **add operations without modifying element classes**
- You want to **collect related behavior externally**

---

## Common Real-World Examples

- AST traversal in compilers (syntax tree nodes)
- Financial calculations on different account types
- Exporting objects to different formats (JSON, XML, CSV)
- Logging or auditing object structures

---

## Example in this project

We have a **Shopping Cart**:
- Items: Book, Fruit
- Visitors: PriceCalculator, TaxCalculator
- Visitor computes different operations on items
