# State Pattern (Go)

## What is State?

State is a behavioral design pattern.
It allows an object to **alter its behavior when its internal state changes**.
The object will appear to **change its class dynamically**.

You use this pattern when:
- An object has **different behaviors depending on state**
- You want to **avoid large if/else or switch statements**
- You want **state-specific logic encapsulated**

---

## Common Real-World Examples

- Media player (Playing, Paused, Stopped)
- TCP connection states (Closed, Listening, Established)
- Vending machine states
- Document workflow (Draft, Moderation, Published)

---

## Example in this project

We have a **MediaPlayer**:
- States: Playing, Paused, Stopped
- Each state defines its own behavior for Play, Pause, Stop
