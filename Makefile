all: clean build run

ifneq ($(OS),Windows_NT)
    OS := $(shell sh -c 'uname -s 2>/dev/null')
endif

ifeq ($(OS),Linux)
    LD_FLAGS = -ldflags="-s -w"
endif

.PHONY: build
build:
	CGO_ENABLED=0 go build $(LD_FLAGS) -o mdtosnow cmd/mdtosnow/main.go

.PHONY: clean
clean:
	rm mdtosnow

.PHONY: deps-init
deps-init:
	rm -f go.mod go.sum
	go mod init
	go mod tidy

.PHONY: deps
deps:
	go mod download

.PHONY: lint
lint:
	golangci-lint run

.PHONY: run
run:
	./mdtosnow
