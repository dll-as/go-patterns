# MVP (Model-View-Presenter) in Go

## What is MVP?

- **Model:** data and business logic
- **View:** responsible only for UI / output
- **Presenter:** handles user input, updates Model and View

Use when:
- Building UI apps (desktop/web)
- Want testable presenters
- Keep View dumb and reusable
