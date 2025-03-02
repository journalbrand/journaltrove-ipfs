<!-- === WATCHER HEADER START === -->
<!-- File: todo-ipfs/README.md -->
<!-- Managed by file watcher -->
<!-- === WATCHER HEADER END === -->
# 🌐 Todo App IPFS Node

This repository contains the IPFS node implementation for the Todo App ecosystem, providing decentralized storage capabilities for todo lists.

## 📚 Overview

The IPFS node serves as the backbone for secure, distributed storage in the Todo App ecosystem. It enables:

- Decentralized storage of todo lists
- Cross-device synchronization
- Cryptographic verification of content
- Peer-to-peer distribution of data
- Offline-first functionality

## 🧩 Current Implementation

The IPFS node currently implements:

- ✅ Basic Go application structure
- ✅ Echo Service that satisfies initial requirements (System.1.1.IPFS.1)
- ✅ Unit tests with requirement mappings
- ✅ CI integration with the system repository
- 🚧 IPFS core functionality integration (in progress)
- 🚧 Todo data schema and validation (planned)
- 🚧 Identity management (planned)

## 🔧 Technology Stack

- **Language**: Go 1.19+
- **IPFS Implementation**: go-ipfs / kubo
- **Testing Framework**: Go testing package
- **Build System**: Go modules with Make
- **CI Pipeline**: GitHub Actions

## 🚀 Getting Started

### Prerequisites

- Go 1.19 or newer
- Git
- Make
- IPFS knowledge (recommended)

### Development Setup

1. Clone this repository:
```bash
git clone https://github.com/journalbrand/todo-ipfs.git
```

2. Install dependencies:
```bash
cd todo-ipfs
go mod download
```

3. Build the application:
```bash
make build
```

4. Run tests:
```bash
make test
```

### Project Structure

```
todo-ipfs/
├── cmd/                  # Command-line applications
│   └── todo-ipfs/        # Main IPFS node application
│       └── main.go       # Application entry point
├── pkg/                  # Reusable packages
│   └── echo/             # Echo service implementation
│       ├── echo.go       # Echo service code
│       └── echo_test.go  # Echo service tests
├── tests/                # Test configurations
│   └── test-mappings.jsonld # Mapping of tests to requirements
├── requirements/         # Component requirements
│   └── requirements.jsonld # IPFS-specific requirements
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
└── Makefile             # Build automation
```

## 📋 Requirements

The IPFS node implements specific requirements defined in the JSON-LD format:

- Requirements are stored in `requirements/requirements.jsonld`
- Tests are mapped to requirements in `tests/test-mappings.jsonld`
- Every implementation must satisfy specific requirements

### Key Requirements

- **System.1.1.IPFS.1**: Input handling capability (implemented in Echo service)
- **System.2.1.IPFS.1**: IPFS content storage (planned)
- **System.2.1.IPFS.2**: IPFS content retrieval (planned)
- **System.2.1.IPFS.3**: IPFS node discovery (planned)
- **System.3.1.IPFS.1**: Cryptographic identity (planned)

## 🧪 Testing

All code must be thoroughly tested and mapped to requirements:

1. Each test should verify a specific requirement
2. Test mapping should be defined in `tests/test-mappings.jsonld`
3. Run tests before submitting changes:
```bash
make test
```

## 🔄 CI/CD Integration

This repository integrates with the Todo App CI/CD pipeline:

- **CI Workflow**: `.github/workflows/ci.yml`
- **Artifacts**: Test results in JSON-LD format
- **Integration**: Test results are sent to the system repository for compliance checks

## 🔌 IPFS Integration

The Todo App IPFS node will utilize several IPFS capabilities:

- **Content Addressing**: Data is referenced by its content hash, ensuring integrity
- **Distributed Hash Table (DHT)**: For peer and content discovery
- **Merkle DAG**: For efficient data structure representation
- **PubSub**: For real-time updates between clients
- **IPNS**: For mutable pointers to todo lists

## 🔗 Related Repositories

The IPFS node is part of a multi-repository ecosystem:

- [todo-system](https://github.com/journalbrand/todo-system) - System-level coordination, requirements and CI/CD orchestration
- [todo-ios](https://github.com/journalbrand/todo-ios) - iOS client
- [todo-android](https://github.com/journalbrand/todo-android) - Android client

## 📝 Contributing

To contribute to the IPFS node:

1. Fork the repository
2. Create a feature branch
3. Implement your changes following Go best practices
4. Add tests that map to requirements
5. Ensure the CI workflow passes
6. Submit a pull request

## 📜 License

This project is licensed under the MIT License - see the LICENSE file for details. 
