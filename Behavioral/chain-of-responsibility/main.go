package main

import "fmt"

// Handler interface
type Logger interface {
	SetNext(next Logger)
	Log(level string, message string)
}

// Base handler
type BaseLogger struct {
	next Logger
}

func (b *BaseLogger) SetNext(next Logger) {
	b.next = next
}

// Concrete handlers
type InfoLogger struct {
	BaseLogger
}

func (i *InfoLogger) Log(level string, message string) {
	if level == "INFO" {
		fmt.Println("INFO:", message)
	} else if i.next != nil {
		i.next.Log(level, message)
	}
}

type DebugLogger struct {
	BaseLogger
}

func (d *DebugLogger) Log(level string, message string) {
	if level == "DEBUG" {
		fmt.Println("DEBUG:", message)
	} else if d.next != nil {
		d.next.Log(level, message)
	}
}

type ErrorLogger struct {
	BaseLogger
}

func (e *ErrorLogger) Log(level string, message string) {
	if level == "ERROR" {
		fmt.Println("ERROR:", message)
	} else if e.next != nil {
		e.next.Log(level, message)
	}
}

func main() {
	// Setup chain: Info -> Debug -> Error
	info := &InfoLogger{}
	debug := &DebugLogger{}
	errorLog := &ErrorLogger{}

	info.SetNext(debug)
	debug.SetNext(errorLog)

	// Test logging
	info.Log("INFO", "This is an info message")
	info.Log("DEBUG", "This is a debug message")
	info.Log("ERROR", "This is an error message")
}
