# Configuration via Struct + Tags (Go Idiomatic Pattern)

## What is Configuration via Struct + Tags?

- Define a struct for your configuration
- Use tags to specify source (json, env, default)
- Parse config into struct at startup

Use when:
- Your app has many configurable parameters
- You want typed and structured configuration
- You want to load config from files, environment, or flags
