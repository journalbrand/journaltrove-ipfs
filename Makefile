# === WATCHER HEADER START ===
# File: todo-ipfs/Makefile
# Managed by file watcher
# === WATCHER HEADER END ===
.PHONY: build test clean run

build:
	go build -o todo_app

test:
	go test ./... -v

clean:
	rm -f todo_app

run: build
	./todo_app

lint:
	go vet ./...

all: clean build test 
