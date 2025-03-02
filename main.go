// === WATCHER HEADER START ===
// File: journaltrove-ipfs/main.go
// Managed by file watcher
// === WATCHER HEADER END ===
package main

import (
	"fmt"
	"github.com/journalbrand/journaltrove-ipfs/pkg/echo"
)

func main() {
	service := echo.NewService()
	
	message := "Hello from IPFS journaltrove App!"
	
	// Simple echo
	result := service.Echo(message)
	fmt.Println(result)
	
	// Echo with timestamp
	timestampResult := service.EchoWithTimestamp(message)
	fmt.Println(timestampResult)
} 
