# Decorator Pattern (Go)

## What is Decorator?

Decorator is a structural design pattern.
It allows you to **add new behavior to objects dynamically** without modifying their code.

You use this pattern when:
- You want to add responsibilities at runtime
- Subclassing is impractical or creates too many classes
- You want flexible and reusable enhancements

---

## Common Real-World Examples

- Logging / Metrics / Tracing
- Authentication / Authorization
- Middleware in HTTP servers
- UI enhancements (scrollbars, borders, shadows)

---

## Example in this project

We have a **Notifier** interface:
- Base notifier sends messages
- Decorators add features like logging or encryption



عالی، حالا می‌رسیم به **Decorator Pattern** از **Structural Patterns**.
Decorator معمولاً وقتی استفاده می‌شود که بخواهیم **رفتار یک آبجکت را به‌صورت داینامیک گسترش دهیم بدون تغییر کلاس اصلی**.
مثال مرسوم: **Logging یا Authentication wrapper برای Service**.

---

## 📁 Folder Structure

```text
decorator/
├── README.md
└── main.go
```

---

## 📄 `README.md`

```md
# Decorator Pattern (Go)

## What is Decorator?

Decorator is a structural design pattern.
It allows you to **add new behavior to objects dynamically** without modifying their code.

You use this pattern when:
- You want to add responsibilities at runtime
- Subclassing is impractical or creates too many classes
- You want flexible and reusable enhancements

---

## Common Real-World Examples

- Logging / Metrics / Tracing
- Authentication / Authorization
- Middleware in HTTP servers
- UI enhancements (scrollbars, borders, shadows)

---

## Example in this project

We have a **Notifier** interface:
- Base notifier sends messages
- Decorators add features like logging or encryption
```

## 🧠 Key Points (Simple)

* Base component implements `Notifier`
* Decorators wrap the component and implement the same interface
* Can chain multiple decorators dynamically
* No need to modify the base component

---

## 🆚 Decorator vs Adapter

| Decorator                    | Adapter                           |
| ---------------------------- | --------------------------------- |
| Adds behavior                | Changes interface                 |
| Wrap object dynamically      | Make incompatible interfaces work |
| Example: Logging / UpperCase | Example: StripeAdapter            |