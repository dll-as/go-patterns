package main

import "fmt"

// Base struct
type Logger struct{}

func (l *Logger) Log(msg string) {
	fmt.Println("[LOG]:", msg)
}

// Embedded struct (composition)
type Service struct {
	Logger      // embedding Logger
	ServiceName string
}

func (s *Service) Start() {
	s.Log("Starting service: " + s.ServiceName)
}

func main() {
	svc := Service{ServiceName: "MyService"}
	svc.Start() // can directly call Log() from embedded Logger
}
