// Package echo provides a simple string echo service
package echo

import (
	"fmt"
	"time"
)

// Service represents an echo service
type Service struct{}

// NewService creates a new echo service
func NewService() *Service {
	return &Service{}
}

// Echo simply returns the provided message
// This implements Req.IPFS.1 by demonstrating input handling capability
func (s *Service) Echo(message string) string {
	return message
}

// EchoWithTimestamp returns the message with a UTC timestamp appended
func (s *Service) EchoWithTimestamp(message string) string {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	return fmt.Sprintf("%s (timestamp: %s)", message, timestamp)
}
