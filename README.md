# Go Design Patterns Collection

A comprehensive collection of **Go Design Patterns**, including:

- Creational Patterns
- Structural Patterns
- Behavioral Patterns
- Concurrency Patterns
- Go-specific / Idiomatic Patterns
- Architectural / Enterprise Patterns

This repository is designed for learning, reference, and practical examples in Go.
Each pattern has its own directory with:

- `main.go` → example code
- `README.md` → explanation & usage


## 📚 Table of Contents

### Creational Patterns

- [Singleton](./Creational/singleton)
- [Factory Method](./Creational/factory-method)
- [Abstract Factory](./Creational/abstract-factory)
- [Builder](./Creational/builder)
- [Prototype](./Creational/prototype)

### Structural Patterns

- [Adapter](./Structural/adapter)
- [Bridge](./Structural/bridge)
- [Composite](./Structural/composite)
- [Decorator](./Structural/decorator)
- [Facade](./Structural/facade)
- [Flyweight](./Structural/flyweight)
- [Proxy](./Structural/proxy)

### Behavioral Patterns

- [Chain of Responsibility](./Behavioral/chain-of-responsibility)
- [Command](./Behavioral/command)
- [Interpreter](./Behavioral/interpreter)
- [Iterator](./Behavioral/iterator)
- [Mediator](./Behavioral/mediator)
- [Memento](./Behavioral/memento)
- [Observer](./Behavioral/observer)
- [State](./Behavioral/state)
- [Strategy](./Behavioral/strategy)
- [Template Method](./Behavioral/template-method)
- [Visitor](./Behavioral/visitor)

### Concurrency Patterns

- [Goroutine per Task](./Concurrency/goroutine-per-task)
- [Worker Pool](./Concurrency/worker-pool)
- [Fan-In](./Concurrency/fan-in)
- [Fan-Out](./Concurrency/fan-out)
- [Pipeline](./Concurrency/pipeline)
- [Future / Promise](./Concurrency/future-promise)
- [Barrier](./Concurrency/barrier)
- [Mutex](./Concurrency/mutex)
- [Semaphore](./Concurrency/semaphore)
- [Read-Write Lock](./Concurrency/rwmutex)
- [Atomic Operations](./Concurrency/atomic)
- [Context Cancellation](./Concurrency/context-cancellation)

### Go-Specific / Idiomatic Patterns

- [Functional Options](./Go-Specific/functional-options)
- [Interface Segregation (Go-style)](./Go-Specific/interface-segregation)
- [Dependency Injection (Constructor-based)](./Go-Specific/dependency-injection)
- [Error Wrapping](./Go-Specific/error-wrapping)
- [Sentinel Errors](./Go-Specific/sentinel-errors)
- [Error Type Assertion](./Go-Specific/error-type-assertion)
- [Empty Interface Pattern](./Go-Specific/empty-interface)
- [Type Assertion Pattern](./Go-Specific/type-assertion)
- [Embedding (Composition over Inheritance)](./Go-Specific/embedding)
- [Table-Driven Tests](./Go-Specific/table-driven-tests)
- [Graceful Shutdown](./Go-Specific/graceful-shutdown)
- [Configuration via Struct + Tags](./Go-Specific/configuration)

### Architectural / Enterprise Patterns

- [MVC](./Architectural/MVC)
- [MVP](./Architectural/MVP)
- [MVVM](./Architectural/MVVM)
- [Clean Architecture](./Architectural/clean-architecture)
- [Hexagonal Architecture (Ports & Adapters)](./Architectural/hexagonal-architecture)
- [Layered Architecture](./Architectural/layered-architecture)
- [Repository Pattern](./Architectural/repository-pattern)
- [Service Layer](./Architectural/service-layer)
- [CQRS](./Architectural/CQRS)
- [Event-Driven Architecture](./Architectural/event-driven-architecture)

## 📖 How to Use

1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/go-design-patterns.git
   cd go-design-patterns
   ```

2. Explore any pattern:
   ```bash
   cd Creational/singleton
   go run main.go
   ```

3. Read the `README.md` in each directory for explanation and examples.


## ⚡ Contribution

Feel free to contribute new patterns, improve examples, or fix typos.
PRs are welcome!
