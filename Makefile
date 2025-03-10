# === WATCHER HEADER START ===
# File: journaltrove-ipfs/Makefile
# Managed by file watcher
# === WATCHER HEADER END ===

.PHONY: build test clean run

build:
	go build -o journaltrove_app

test:
	go test ./... -v

clean:
	rm -f journaltrove_app

run: build
	./journaltrove_app

lint:
	go vet ./...

all: clean build test 
