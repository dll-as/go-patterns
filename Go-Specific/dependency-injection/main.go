package main

import "fmt"

// Logger interface
type Logger interface {
	Log(message string)
}

// Concrete logger
type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("[LOG]:", message)
}

// Service depends on Logger
type Service struct {
	logger Logger
}

// Constructor-based Dependency Injection
func NewService(logger Logger) *Service {
	return &Service{logger: logger}
}

func (s *Service) DoWork() {
	s.logger.Log("Service is doing work")
}

func main() {
	// Inject dependency
	logger := &ConsoleLogger{}
	service := NewService(logger)

	service.DoWork()
}
