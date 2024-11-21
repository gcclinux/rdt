package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

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
