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
	input := "hello"

	// When
	output := service.Echo(input)

	// Then
	if output != input {
		t.Errorf("Echo(%q) = %q, want %q", input, output, input)
	}
}

// TestEchoWithTimestamp verifies that the EchoWithTimestamp method includes both the
// original message and a timestamp
func TestEchoWithTimestamp(t *testing.T) {
	// Given
	service := NewService()
	input := "hello"

	// When
	output := service.EchoWithTimestamp(input)

	// Then
	if !strings.Contains(output, input) {
		t.Errorf("EchoWithTimestamp(%q) = %q, which doesn't contain the input string", input, output)
	}

	// Check if timestamp might be included
	if len(output) <= len(input) {
		t.Errorf("EchoWithTimestamp(%q) = %q, which doesn't appear to include a timestamp", input, output)
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
		t.Errorf("Echo(%q) = %q, want %q", input, output, input)
	}
}

// TestIdentityVerificationFoo verifies that the IdentityVerificationFoo method
// returns the string "foo" as specified by System.2.1.IPFS.5.
func TestIdentityVerificationFoo(t *testing.T) {
	svc := NewService()
	result := svc.IdentityVerificationFoo()
	expected := "foo"
	if result != expected {
		t.Errorf("IdentityVerificationFoo() = %q, want %q", result, expected)
	}
}

// Added comment to trigger git
