# Adapter Pattern (Go)

## What is Adapter?

Adapter is a structural design pattern.
It allows objects with **incompatible interfaces** to work together.

You use this pattern when:
- You want to use a library or service without modifying it
- Your code expects a different interface
- You want to integrate third-party APIs seamlessly

---

## Common Real-World Examples

- Payment gateways (Stripe, PayPal)
- Logging libraries (Zap, Logrus)
- External APIs (REST / gRPC)
- Legacy code integration

---

## Example in this project

We have a **Payment interface** in our app:
- `Pay(amount float64)`

We want to integrate **Stripe** which has a different method name.
Adapter makes them compatible.

---

## 🧠 Key Points (Simple)

* Adapter wraps an existing object
* Makes it compatible with your interface
* Client code uses only the interface
* Easy to switch or add new adapters later

---

## 🆚 Adapter vs Decorator

| Adapter                 | Decorator                  |
| ----------------------- | -------------------------- |
| Change interface        | Add behavior               |
| Compatible with client  | Wrap object behavior       |
| Example: Stripe adapter | Example: Logging decorator |
