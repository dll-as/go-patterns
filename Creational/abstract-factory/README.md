# Abstract Factory Pattern (Go)

## What is Abstract Factory?

Abstract Factory is a creational design pattern.
It provides an interface to create **families of related objects**
without specifying their concrete types.

You use this pattern when:
- You have multiple product types that must work together
- You want to switch entire families of objects at once
- You want to avoid mixing incompatible objects

---

## Common Real-World Examples

- Database systems (MySQL, PostgreSQL)
  - Connection
  - Query builder
  - Transaction
- UI toolkits (Windows / Mac)
  - Button
  - Checkbox
- Cloud providers (AWS / GCP)
  - Storage
  - Queue
  - Logger

---

## Example in this project

We support **two databases**:
- MySQL
- PostgreSQL

Each database provides:
- Connection
- Query builder

The factory creates matching components together.