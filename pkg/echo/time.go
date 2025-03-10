package echo

import (
	"time"
)

// CurrentTime returns the current system time
// Implements requirement System.4.IPFS.1
func CurrentTime() time.Time {
	return time.Now()
}

// CurrentTimeString returns the current time as a formatted string
// Implements requirement System.4.IPFS.1
func CurrentTimeString() string {
	return time.Now().Format(time.RFC1123)
} 