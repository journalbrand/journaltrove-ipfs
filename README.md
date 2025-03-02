# Todo IPFS Implementation

This repository contains the Go implementation of the Todo App IPFS component.

## Current Status

This implementation currently provides a simple Echo service that demonstrates input handling capabilities, satisfying requirement `Req.IPFS.1`.

## Components

- **Echo Service**: A simple service that takes an input string and echoes it back, with an additional method that includes a timestamp.

## Getting Started

### Prerequisites

- Go 1.19 or higher

### Building the Application

```bash
go build -o todo_app
# Or use the Make target
make build
```

### Running the Application

```bash
./todo_app
# Or use the Make target
make run
```

### Running Tests

```bash
go test ./...
# Or use the Make target
make test
```

## Project Structure

```
todo-ipfs/
├── go.mod          # Go module definition
├── main.go         # Main application entry point
├── Makefile        # Build automation
└── pkg/
    └── echo/       # Echo service implementation
        ├── echo.go      # Service implementation
        └── echo_test.go # Unit tests
```

## Requirements Mapping

- `Req.IPFS.1`: Input handling capability - Implemented in the Echo service

## 🚀 Current State

This repository is currently in the requirements definition phase. Implementation will begin soon.

## 📋 Requirements

Requirements are defined in JSON-LD format in the `requirements/requirements.jsonld` file.

## 🔗 Related Repositories

- [todo-system](https://github.com/journalbrand/todo-system) - System-level requirements and coordination
- [todo-ios](https://github.com/journalbrand/todo-ios) - iOS client
- [todo-android](https://github.com/journalbrand/todo-android) - Android client

## 🧪 Development

Development will follow the implementation plan outlined in the system repository. 