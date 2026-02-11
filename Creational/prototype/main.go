package main

import "fmt"

// Prototype interface
type Cloneable interface {
	Clone() Cloneable
}

// Concrete prototype
type ServerConfig struct {
	Host    string
	Port    int
	Timeout int
}

// Clone creates a copy of ServerConfig
func (s *ServerConfig) Clone() Cloneable {
	copy := *s // shallow copy
	return &copy
}

func main() {
	// base config (expensive to create)
	baseConfig := &ServerConfig{
		Host:    "localhost",
		Port:    8080,
		Timeout: 30,
	}

	// clone and customize
	devConfig := baseConfig.Clone().(*ServerConfig)
	devConfig.Port = 3000

	prodConfig := baseConfig.Clone().(*ServerConfig)
	prodConfig.Host = "prod.server.com"
	prodConfig.Timeout = 60

	fmt.Println("Base:", baseConfig)
	fmt.Println("Dev:", devConfig)
	fmt.Println("Prod:", prodConfig)
}
