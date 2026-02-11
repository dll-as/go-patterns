package main

import (
	"encoding/json"
	"fmt"
)

// Config struct with tags
type Config struct {
	Port     int    `json:"port"`
	LogLevel string `json:"log_level"`
	Debug    bool   `json:"debug"`
}

func main() {
	// Example JSON configuration file
	configJSON := `{"port":8080, "log_level":"info", "debug":true}`

	var cfg Config
	if err := json.Unmarshal([]byte(configJSON), &cfg); err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	fmt.Printf("Loaded Config: %+v\n", cfg)

	// Use config
	if cfg.Debug {
		fmt.Println("Debug mode enabled")
	}
	fmt.Printf("Server will run on port %d with log level %s\n", cfg.Port, cfg.LogLevel)
}
