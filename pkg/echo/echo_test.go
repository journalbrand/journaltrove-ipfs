// === WATCHER HEADER START ===
// File: journaltrove-ipfs/pkg/echo/echo_test.go
// Managed by file watcher
// === WATCHER HEADER END ===
package echo

import (
	"strings"
	"testing"
)

// TestEcho verifies that the Echo method returns the exact input string
func TestEcho(t *testing.T) {
	// Given
	service := NewService()
	input := "Hello, World!"

	// When
	output := service.Echo(input)

	// Then
	if output != input {
		t.Errorf("Echo should return the same string that was input: got %v want %v", output, input)
	}
}

// TestEchoWithTimestamp verifies that the EchoWithTimestamp method includes both the
// original message and a timestamp
func TestEchoWithTimestamp(t *testing.T) {
	// Given
	service := NewService()
	input := "Hello, World!"

	// When
	output := service.EchoWithTimestamp(input)

	// Then
	if !strings.Contains(output, input) {
		t.Errorf("Output should contain the original input: %v", output)
	}

	if !strings.Contains(output, "timestamp:") {
		t.Errorf("Output should contain a timestamp: %v", output)
	}
}

// TestEchoWithEmptyString verifies that the Echo method works with empty strings
func TestEchoWithEmptyString(t *testing.T) {
	// Given
	service := NewService()
	input := ""

	// When
	output := service.Echo(input)

	// Then
	if output != input {
		t.Errorf("Echo should work with empty strings: got %v want %v", output, input)
	}
}

// TestIdentityVerificationFoo verifies that the IdentityVerificationFoo method
// returns the string "foo" as required by System.3.IPFS.1
func TestIdentityVerificationFoo(t *testing.T) {
	svc := NewService()
	result := svc.IdentityVerificationFoo()
	if result != "foo" {
		t.Errorf("IdentityVerificationFoo() failed, expected: foo, got: %s", result)
	}
}

// Added comment to trigger git
