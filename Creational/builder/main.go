package main

import "fmt"

// Product
type Server struct {
	Port    int
	Timeout int
	TLS     bool
	Logging bool
}

// Builder
type ServerBuilder struct {
	server Server
}

// NewServerBuilder creates a new builder with default values
func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{
		server: Server{
			Port:    8080,
			Timeout: 30,
			TLS:     false,
			Logging: false,
		},
	}
}

func (b *ServerBuilder) SetPort(port int) *ServerBuilder {
	b.server.Port = port
	return b
}

func (b *ServerBuilder) SetTimeout(timeout int) *ServerBuilder {
	b.server.Timeout = timeout
	return b
}

func (b *ServerBuilder) EnableTLS() *ServerBuilder {
	b.server.TLS = true
	return b
}

func (b *ServerBuilder) EnableLogging() *ServerBuilder {
	b.server.Logging = true
	return b
}

// Build returns the final product
func (b *ServerBuilder) Build() Server {
	return b.server
}

func main() {
	server := NewServerBuilder().
		SetPort(9090).
		SetTimeout(60).
		EnableTLS().
		EnableLogging().
		Build()

	fmt.Printf("%+v\n", server)
}
