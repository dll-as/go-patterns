<p align="center">
  <img src="./assets/go-logo.png" width="200">
  <h1 align="center">🐹 Go Design Patterns Collection</h1>
  <p align="center">
    A comprehensive, practical, and production-ready collection of Go design patterns
    <br/>
    <strong>Learn • Reference • Practice</strong>
  </p>
  <p align="center">
    <img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat&logo=go">
    <img src="https://img.shields.io/badge/Patterns-50+-success">
    <img src="https://img.shields.io/badge/License-MIT-blue">
    <img src="https://img.shields.io/badge/PRs-Welcome-brightgreen">
  </p>
</p>


## 🌟 Why This Repository?

This isn't just another list of Go patterns. It's a **hands-on, production-aware** collection where every pattern comes with:

| Feature | Description |
|---------|-------------|
| ✅ **Runable Code** | Every pattern has a `main.go` you can execute immediately |
| 📖 **Detailed README** | Each directory explains when/why/how to use the pattern |
| 🎯 **Go-Specific** | Patterns adapted to Go idioms, not Java/C++ translations |
| 🔧 **Practical Examples** | Real-world scenarios, not just theoretical `Animal`/`Shape` examples |
| 📊 **Multiple Categories** | From classic Gang of Four to modern Go concurrency patterns |


## 📚 Table of Contents

<details>
<summary><b>🏗️ Creational Patterns</b> — Object creation mechanisms</summary>

- [Singleton](./Creational/singleton) — Ensure a type has only one instance
- [Factory Method](./Creational/factory-method) — Delegate object creation to subclasses
- [Abstract Factory](./Creational/abstract-factory) — Create families of related objects
- [Builder](./Creational/builder) — Construct complex objects step by step
- [Prototype](./Creational/prototype) — Clone objects instead of creating from scratch
</details>

<details>
<summary><b>🔌 Structural Patterns</b> — Object composition and relationships</summary>

- [Adapter](./Structural/adapter) — Make incompatible interfaces compatible
- [Bridge](./Structural/bridge) — Separate abstraction from implementation
- [Composite](./Structural/composite) — Treat individual and composite objects uniformly
- [Decorator](./Structural/decorator) — Add behavior without inheritance
- [Facade](./Structural/facade) — Simplify complex subsystems
- [Flyweight](./Structural/flyweight) — Share objects to support large quantities
- [Proxy](./Structural/proxy) — Control access to another object
</details>

<details>
<summary><b>🔄 Behavioral Patterns</b> — Object communication and responsibility</summary>

- [Chain of Responsibility](./Behavioral/chain-of-responsibility) — Pass requests along a chain
- [Command](./Behavioral/command) — Encapsulate requests as objects
- [Interpreter](./Behavioral/interpreter) — Define a grammar and interpret sentences
- [Iterator](./Behavioral/iterator) — Traverse collections without exposing internals
- [Mediator](./Behavioral/mediator) — Reduce coupling between objects
- [Memento](./Behavioral/memento) — Capture and restore object state
- [Observer](./Behavioral/observer) — Notify dependents of state changes
- [State](./Behavioral/state) — Alter behavior when state changes
- [Strategy](./Behavioral/strategy) — Select algorithms at runtime
- [Template Method](./Behavioral/template-method) — Define skeleton of an algorithm
- [Visitor](./Behavioral/visitor) — Separate algorithms from objects
</details>

<details>
<summary><b>⚡ Concurrency Patterns</b> — Go's superpower! Goroutines & channels</summary>

- [Goroutine per Task](./Concurrency/goroutine-per-task) — Simple concurrent execution
- [Worker Pool](./Concurrency/worker-pool) — Control concurrent goroutines
- [Fan-In](./Concurrency/fan-in) — Multiplex multiple channels into one
- [Fan-Out](./Concurrency/fan-out) — Distribute work across goroutines
- [Pipeline](./Concurrency/pipeline) — Stage-wise data processing
- [Future/Promise](./Concurrency/future-promise) — Async result handling
- [Barrier](./Concurrency/barrier) — Synchronize multiple goroutines
- [Mutex](./Concurrency/mutex) — Mutual exclusion
- [Semaphore](./Concurrency/semaphore) — Limit concurrent access
- [RWMutex](./Concurrency/rwmutex) — Reader-writer locks
- [Atomic](./Concurrency/atomic) — Lock-free operations
- [Context Cancellation](./Concurrency/context-cancellation) — Graceful cancellation
</details>

<details>
<summary><b>🎯 Go-Specific Patterns</b> — Idiomatic Go practices</summary>

- [Functional Options](./Go-Specific/functional-options) — Clean, extensible constructors
- [Interface Segregation](./Go-Specific/interface-segregation) — Small, focused interfaces
- [Dependency Injection](./Go-Specific/dependency-injection) — Constructor-based DI
- [Error Wrapping](./Go-Specific/error-wrapping) — Contextual error handling (Go 1.13+)
- [Sentinel Errors](./Go-Specific/sentinel-errors) — Predefined error variables
- [Error Type Assertion](./Go-Specific/error-type-assertion) — Custom error types
- [Empty Interface](./Go-Specific/empty-interface) — `interface{}` and `any` usage
- [Type Assertion](./Go-Specific/type-assertion) — Safe type conversion
- [Embedding](./Go-Specific/embedding) — Composition over inheritance
- [Table-Driven Tests](./Go-Specific/table-driven-tests) — Clean, maintainable tests
- [Graceful Shutdown](./Go-Specific/graceful-shutdown) — Clean server termination
- [Configuration](./Go-Specific/configuration) — Struct tags and config loading
</details>

<details>
<summary><b>🏛️ Architectural Patterns</b> — System-level organization</summary>

- [MVC](./Architectural/MVC) — Model-View-Controller
- [MVP](./Architectural/MVP) — Model-View-Presenter
- [MVVM](./Architectural/MVVM) — Model-View-ViewModel
- [Clean Architecture](./Architectural/clean-architecture) — Uncle Bob's architecture
- [Hexagonal Architecture](./Architectural/hexagonal-architecture) — Ports & Adapters
- [Layered Architecture](./Architectural/layered-architecture) — N-tier architecture
- [Repository Pattern](./Architectural/repository-pattern) — Data access abstraction
- [Service Layer](./Architectural/service-layer) — Business logic boundary
- [CQRS](./Architectural/CQRS) — Command Query Responsibility Segregation
- [Event-Driven Architecture](./Architectural/event-driven-architecture) — Event-based communication
</details>

## 🚀 Quick Start

```bash
# Clone the repository
git clone https://github.com/dll-as/go-patterns.git
cd go-patterns

# Pick a pattern and run it!
cd Creational/singleton
go run main.go

# Or run tests if available
cd Go-Specific/table-driven-tests
go test -v
```

## 📖 How to Read This Repository
   - ### 🎓 **For Beginners**
      Start with **Creational** patterns — they're the most intuitive. Then move to **Structural** and **Behavioral**. Each pattern's README explains the "what" and "why" before the "how".

   - ### 🚗 **For Experienced Gophers**
      Jump straight to **Concurrency** and **Go-Specific** patterns. These show you how Go differs from other launguages and how to write idiomatic Go code.

   - ### 🏗️ **For Architects**
      The **Architectural** patterns section demonstrates how to structure entire applications. See how clean architecture and hexagonal architecture are implemented in Go.

   - ## ✨ What Makes a Good Pattern Example?

      Each pattern directory contains:

      ```
      pattern-name/
      ├── main.go           # Runnable, self-contained example
      ├── README.md         # Detailed explanation with:
      │                     # • Intent & Problem
      │                     # • Solution & Structure
      │                     # • When to use/avoid
      │                     # • Real-world examples
      │                     # • Comparison with other patterns
      └── main_test.go      # Optional: Tests showing usage
      ```