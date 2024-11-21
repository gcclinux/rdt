package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// handleConnection is a function that handles a new connection
func handleConnection(inbound net.Conn, config *Config) {
	// Create log files
	inboundLog, err := createLogFile(config, "inbound")
	if err != nil {
		log.Printf("Failed to create inbound log: %v", err)
		return
	}
	defer inboundLog.Close()

	outboundLog, err := createLogFile(config, "outbound")
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
