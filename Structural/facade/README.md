# Facade Pattern (Go)

## What is Facade?

Facade is a structural design pattern.
It provides a **simplified interface** to a complex subsystem.

You use this pattern when:
- The subsystem is complicated
- You want to hide internal complexity
- You want a single entry point for clients

---

## Common Real-World Examples

- Unified API for multiple services (Auth, Payment, Notification)
- Database access layer
- Multimedia libraries
- Cloud SDK wrappers

---

## Example in this project

We have three subsystems:
1. AuthenticationService
2. PaymentService
3. NotificationService

Facade provides a **simple method** for the client to execute all steps.

---

## 🧠 Key Points (Simple)

* Facade simplifies client interaction
* Client does not need to know internal subsystem details
* Easy to extend subsystem without affecting client code

---

## 🆚 Facade vs Adapter

| Facade                    | Adapter                                |
| ------------------------- | -------------------------------------- |
| Simplify complex system   | Make incompatible interface compatible |
| Hides internal complexity | Converts interface                     |
| Example: ShopFacade       | Example: StripeAdapter                 |