BIN=bin/zebrapad

.PHONY: all build test run clean

all: clean test build run

$(BIN): build
build: src/*.go src/web/*
	mkdir -p bin
	cd src && go build -o ../$(BIN) .

test:
	cd src && go test .

run: $(BIN)
	./$(BIN)

clean:
	rm -rf bin/
