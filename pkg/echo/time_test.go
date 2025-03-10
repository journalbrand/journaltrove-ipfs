package echo

import (
	"strings"
	"testing"
	"time"
)

// TestCurrentTime verifies that the current time function returns a time close to now
// This test verifies requirement System.4.IPFS.1
func TestCurrentTime(t *testing.T) {
	// When we get the current time
	result := CurrentTime()
	now := time.Now()

	// Then the time should be within 1 second of now
	diff := now.Sub(result)
	if diff < 0 {
		diff = -diff
	}
	if diff > time.Second {
		t.Errorf("Current time should be within 1 second of system time, got diff of %v", diff)
	}
}

// TestCurrentTimeString verifies that the current time string function returns a properly formatted string
// This test verifies requirement System.4.IPFS.1
func TestCurrentTimeString(t *testing.T) {
	// When we get the current time as a string
	result := CurrentTimeString()

	// Then the result should be a non-empty string
	if result == "" {
		t.Error("Current time string should not be empty")
	}

	// And should contain the current year
	currentYear := time.Now().Format("2006")
	if !strings.Contains(result, currentYear) {
		t.Errorf("Current time string should contain the current year %s, but got %s", currentYear, result)
	}
} 