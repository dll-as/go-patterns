# Iterator Pattern (Go)

## What is Iterator?

Iterator is a behavioral design pattern.
It provides a way to **access elements of a collection sequentially without exposing its underlying representation**.

You use this pattern when:
- You want to **iterate over different collections uniformly**
- You want to **hide collection implementation**
- You want **multiple traversal strategies**

---

## Common Real-World Examples

- Slices, arrays, or maps in Go
- Custom collections (trees, graphs)
- Streams or pipelines
- Paginated API results

---

## Example in this project

We have a **collection of books**:
- `Book` struct
- `BookCollection` implements iterator
