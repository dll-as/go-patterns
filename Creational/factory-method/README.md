## What is Factory Method?
Factory Method is a creational design pattern.
It provides an interface to create objects, but lets the code decide
**which concrete type** to create.


You use this pattern when:
- You have multiple implementations of one behavior
- The exact type should be decided at runtime
- You want to avoid tight coupling with concrete structs


## Common Real-World Examples
- Database connections (MySQL, PostgreSQL, MongoDB)
- Notification systems (Email, SMS, Push)
- Payment providers (Stripe, PayPal)
- Logging systems (File, Console, Remote)


## Example in this project
We have a **Notification** service:
- Email notification
- SMS notification

The factory decides which one to create.
