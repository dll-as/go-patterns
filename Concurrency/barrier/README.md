# Barrier Pattern (Go Concurrency Pattern)

## What is Barrier?

Barrier is a synchronization pattern.

Multiple goroutines must:
- Reach a certain point
- Wait for others
- Continue execution together

It ensures that no goroutine moves forward
until all participants reach the barrier.

---

## When to Use

- Parallel computations that need synchronization
- Simulation steps
- Batch processing
- Multi-phase algorithms

---

## Common Real-World Examples

- Parallel matrix computation
- Multiplayer game tick synchronization
- Distributed system phases
- Scientific simulations

---

## Example in this project

We simulate:
- 5 workers doing preparation work
- All wait at the barrier
- After everyone arrives, they continue together
