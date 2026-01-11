# Variables
BINARY_NAME=ledger-cli
MAIN_PATH=./cmd/main.go
DEFAULT_FILE=./test_data/given.csv
DEFAULT_PERIOD=202201

.PHONY: all build test run clean tidy

all: tidy build

build:
	go build -o $(BINARY_NAME) $(MAIN_PATH)

run:
	go run $(MAIN_PATH) -file $(DEFAULT_FILE) -period $(DEFAULT_PERIOD)

clean:
	rm -f $(BINARY_NAME)

tidy:
	go mod tidy
