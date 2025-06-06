NAME=Tic-Tac-Go
GOCMD=LC_ALL=C go
TIMEOUT=5

# go tools
export PATH := ./bin:$(PATH)
export GO111MODULE := on
export GOPROXY = https://proxy.golang.org,direct

# go source files
SRC = $(shell find . -type f -name "*.go")
# The name of the executable (default is current directory name)
TARGET := $(shell echo $${PWD-`pwd`})

.PHONY: all build setup test cover clean run help

## all: Default target, now is build
all: build

## build: Builds the binary
build:
	@echo "files will be saved in build/"
	mkdir -p build
	@echo "Building..."
	@echo "Building - terminal game; linux..."
	GOOS="linux" $(GOCMD) build -v -o build/tic-tac-go-terminal.bin cmd/terminal-game/tic-tac-go.go
	@echo "Building - terminal game; windows..."
	GOOS="windows" $(GOCMD) build -v -o build/tic-tac-go-terminal.exe cmd/terminal-game/tic-tac-go.go
	@echo "Building - windowed game; linux..."
	CGO_ENABLED="1" GOOS="linux" $(GOCMD) build -v -o build/tic-tac-go-windowed.bin cmd/giu-game/tic-tac-go.go
	@echo "Building - windowed game; windows..."
	CGO_ENABLED="1" GOOS="windows" CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ \
		HOST=x86_64-w64-mingw32 \
		$(GOCMD) build -v -ldflags "-s -w -H=windowsgui -extldflags=-static" \
		-o build/tic-tac-go-windowed.exe cmd/giu-game/tic-tac-go.go

## setup: Runs mod download and generate
setup:
	@echo "Tic-Tac-Go INFO: to cross-platform build windowed version of an app,"
	@echo "Tic-Tac-Go INFO: make sure you've mingw compiller installed."
	@echo "Tic-Tac-Go INFO:"
	@echo "Tic-Tac-Go INFO: For more details check https://github.com/AllenDang/giu"
	@echo "Downloading tools and dependencies..."
	@git submodule update --init --recursive
	@$(GOCMD) get golang.org/x/tools/cmd/stringer
	@$(GOCMD) install golang.org/x/tools/cmd/stringer
	@$(GOCMD) get -d ./...
	@$(GOCMD) mod download -x
	@$(GOCMD) generate -v ./...

## test: Runs the tests with coverage
test:
	@echo "Running tests..."
	@$(GOCMD) test -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt ./... -run . -timeout $(TIMEOUT)m

## cover: Runs all tests and opens the coverage report in the browser
cover: test
	@$(GOCMD) tool cover -html=coverage.txt

## clean: Runs go run
clean:
	@echo "Cleaning..."
	@$(GOCMD) clean

## run: Runs go run
run: build
	@$(GOCMD) run ${TARGET}

## help: Prints this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
