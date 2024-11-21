// config.go
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

type Config struct {
	InboundAddress  string `json:"inbound_address"`
	InboundPort     string `json:"inbound_port"`
	OutboundAddress string `json:"outbound_address"`
	OutboundPort    string `json:"outbound_port"`
	Verbose         bool   `json:"verbose"`
}

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

func setupLogging() error {
	return os.MkdirAll("logs", 0755)
}

func createLogFile(prefix string) (*os.File, error) {
	timestamp := time.Now().Format("20060102150405")
	filename := filepath.Join("logs", fmt.Sprintf("%s_%s.log", prefix, timestamp))
	return os.Create(filename)
}

func logConnection(writer io.Writer, format string, v ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(writer, "%s - %s\n", timestamp, fmt.Sprintf(format, v...))
}

func main() {
	if err := setupLogging(); err != nil {
		log.Fatal("Failed to setup logging:", err)
	}

	// Load configuration
	config, err := loadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Listen on configured address
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.InboundAddress, config.InboundPort))
	if err != nil {
		log.Fatal("Failed to start listener:", err)
	}
	defer listener.Close()

	log.Printf("Proxy listening on %s:%s", config.InboundAddress, config.InboundPort)

	for {
		inbound, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(inbound, config)
	}
}

func handleConnection(inbound net.Conn, config *Config) {
	// Create log files
	inboundLog, err := createLogFile("inbound")
	if err != nil {
		log.Printf("Failed to create inbound log: %v", err)
		return
	}
	defer inboundLog.Close()

	outboundLog, err := createLogFile("outbound")
	if err != nil {
		log.Printf("Failed to create outbound log: %v", err)
		return
	}
	defer outboundLog.Close()

	// Log connection details
	logConnection(inboundLog, "New connection from %s", inbound.RemoteAddr())

	outbound, err := net.Dial("tcp", fmt.Sprintf("%s:%s", config.OutboundAddress, config.OutboundPort))
	if err != nil {
		logConnection(inboundLog, "Failed to connect to target: %v", err)
		inbound.Close()
		return
	}
	defer outbound.Close()
	defer inbound.Close()

	// Modified copy function with logging
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, err := inbound.Read(buffer)
			if err != nil {
				if err != io.EOF {
					logConnection(inboundLog, "Error reading from inbound: %v", err)
				}
				return
			}
			logConnection(inboundLog, "Received: %s", string(buffer[:n]))
			outbound.Write(buffer[:n])
		}
	}()

	buffer := make([]byte, 1024)
	for {
		n, err := outbound.Read(buffer)
		if err != nil {
			if err != io.EOF {
				logConnection(outboundLog, "Error reading from outbound: %v", err)
			}
			return
		}
		logConnection(outboundLog, "Sent: %s", string(buffer[:n]))
		inbound.Write(buffer[:n])
	}
}
