# Strategy Pattern (Go)

## What is Strategy?

Strategy is a behavioral design pattern.
It defines a family of algorithms, **encapsulates each one**, and makes them **interchangeable**.
The client can choose which algorithm to use at runtime.

You use this pattern when:
- You have **multiple ways to perform a task**
- You want **to decouple algorithm from context**
- You want **flexible and reusable behavior**

---

## Common Real-World Examples

- Sorting algorithms (BubbleSort, QuickSort, MergeSort)
- Payment methods (CreditCard, PayPal, Crypto)
- Compression strategies (ZIP, RAR, GZIP)
- Discount strategies in e-commerce

---

## Example in this project

We have a **Payment system**:
- Strategies: CreditCardPayment, PayPalPayment
- Context uses strategy to pay dynamically
