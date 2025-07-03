BIN=zebrapad

.PHONY: all clean test build version run

all: clean test build version run

go.mod:
	go mod init github.com/lvdh/zebrapad

clean:
	rm -rf ./$(BIN)

test: go.mod
	go test -v ./...

$(BIN): build
build: go.mod cmd/zebrapad/*.go internal/*
	go build -o ./$(BIN) ./cmd/zebrapad/*.go

version:
	go version -m ./$(BIN)

run: go.mod $(BIN)
	./$(BIN)
