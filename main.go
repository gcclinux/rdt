package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

// loadConfig reads the configuration from a JSON file
func loadConfig() (*Config, error) {
	file, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// main is the entry point of the application
func main() {
	// Load configuration
	config, err := loadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Setup logging
	if err := setupLogging(config); err != nil {
		log.Fatal("Failed to setup logging:", err)
	}

	// Example usage of createLogFile
	logFile, err := createLogFile(config, "example")
	if err != nil {
		log.Fatal("Failed to create log file:", err)
	}
	if logFile != nil {
		defer logFile.Close()
		logConnection(logFile, "This is a log message")
	}

	// Listen on configured address
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.InboundAddress, config.InboundPort))
	if err != nil {
		log.Fatal("Failed to start listener:", err)
	}
	defer listener.Close()

	log.Printf("RDT listening on %s:%s", config.InboundAddress, config.InboundPort)

	for {
		inbound, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(inbound, config)
	}
}
