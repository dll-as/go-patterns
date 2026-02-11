package main

import (
	"fmt"
	"time"
)

// Server struct
type Server struct {
	Port    int
	Timeout time.Duration
	Debug   bool
}

// Option type
type Option func(*Server)

// Functional options
func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func WithDebug(debug bool) Option {
	return func(s *Server) {
		s.Debug = debug
	}
}

// NewServer constructor
func NewServer(opts ...Option) *Server {
	s := &Server{
		Port:    8080,
		Timeout: 30 * time.Second,
		Debug:   false,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func main() {
	server := NewServer(
		WithPort(9090),
		WithTimeout(60*time.Second),
		WithDebug(true),
	)

	fmt.Printf("Server: %+v\n", server)
}
