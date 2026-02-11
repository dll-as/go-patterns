package main

import "fmt"

// Component interface
type Notifier interface {
	Send(message string)
}

// Concrete component
type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Println("Email sent:", message)
}

// Decorator
type NotifierDecorator struct {
	wrapped Notifier
}

func (d *NotifierDecorator) Send(message string) {
	d.wrapped.Send(message)
}

// Concrete decorator 1: Logging
type LoggingDecorator struct {
	NotifierDecorator
}

func NewLoggingDecorator(notifier Notifier) *LoggingDecorator {
	return &LoggingDecorator{NotifierDecorator{wrapped: notifier}}
}

func (l *LoggingDecorator) Send(message string) {
	fmt.Println("Logging: sending message")
	l.wrapped.Send(message)
}

// Concrete decorator 2: UpperCase message
type UpperCaseDecorator struct {
	NotifierDecorator
}

func NewUpperCaseDecorator(notifier Notifier) *UpperCaseDecorator {
	return &UpperCaseDecorator{NotifierDecorator{wrapped: notifier}}
}

func (u *UpperCaseDecorator) Send(message string) {
	upperMessage := fmt.Sprintf("[UPPER] %s", message)
	u.wrapped.Send(upperMessage)
}

func main() {
	base := &EmailNotifier{}

	// Wrap with logging
	logged := NewLoggingDecorator(base)

	// Wrap with uppercase
	decorated := NewUpperCaseDecorator(logged)

	decorated.Send("Hello World!")
}
