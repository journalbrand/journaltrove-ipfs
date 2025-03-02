// === WATCHER HEADER START ===
// File: todo-ipfs/main.go
// Managed by file watcher
// === WATCHER HEADER END ===
package main

import (
	"fmt"
	"github.com/journalbrand/todo-ipfs/pkg/echo"
)

func main() {
	service := echo.NewService()
	
	message := "Hello from IPFS Todo App!"
	
	// Simple echo
	result := service.Echo(message)
	fmt.Println(result)
	
	// Echo with timestamp
	timestampResult := service.EchoWithTimestamp(message)
	fmt.Println(timestampResult)
} 
