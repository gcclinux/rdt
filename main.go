package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"
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

// setupLogging creates the logs directory if verbose logging is enabled and it doesn't already exist
func setupLogging(config *Config) error {
	if config.Verbose {
		if _, err := os.Stat("logs"); os.IsNotExist(err) {
			if err := os.MkdirAll("logs", 0755); err != nil {
				return err
			}
		}
	}
	return nil
}

// createLogFile creates a new log file with a timestamped filename if verbose logging is enabled
func createLogFile(config *Config, prefix string) (*os.File, error) {
	if config.Verbose {
		timestamp := time.Now().Format("20060102150405")
		filename := filepath.Join("logs", fmt.Sprintf("%s_%s.log", prefix, timestamp))
		return os.Create(filename)
	}
	return nil, nil
}

// logConnection logs a formatted message with a timestamp to the provided writer
func logConnection(writer io.Writer, format string, v ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(writer, "%s - %s\n", timestamp, fmt.Sprintf(format, v...))
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
