# Proxy Pattern (Go)

## What is Proxy?

Proxy is a structural design pattern.
It provides a **surrogate or placeholder** for another object to control access to it.

You use this pattern when:
- You want **lazy initialization**
- You want **access control** (authentication/authorization)
- You want **logging, caching, or monitoring**

---

## Common Real-World Examples

- Virtual Proxy: large images, files loaded on demand
- Protection Proxy: access to sensitive resources
- Remote Proxy: remote services or APIs
- Cache Proxy: cache expensive computations

---

## Example in this project

We have a **Database service**:
- RealDatabase: executes queries
- DatabaseProxy: controls access and logs queries
