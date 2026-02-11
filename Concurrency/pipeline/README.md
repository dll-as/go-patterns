# Pipeline Pattern (Go Concurrency Pattern)

## What is Pipeline?

Pipeline is a concurrency pattern where data flows through multiple stages.
Each stage runs in its own goroutine and communicates using channels.

Each stage:
- Receives input from a channel
- Processes the data
- Sends output to the next stage

You use this pattern when:
- You have multi-step processing
- Each step can run independently
- You want clean and modular concurrency

---

## Common Real-World Examples

- Image processing pipeline
- Data transformation pipeline
- Log processing
- Stream processing systems

---

## Example in this project

We build a simple number pipeline:

1. Generate numbers
2. Square numbers
3. Print results