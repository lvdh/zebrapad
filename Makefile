BIN=zebrapad

.PHONY: all clean test build version run

all: clean test build version run

clean:
	rm -rf ./$(BIN)

go.mod:
	go mod init github.com/lvdh/zebrapad

test: go.mod
	go test ./...

$(BIN): build
build: go.mod cmd/zebrapad/*.go internal/*
	go build -o ./$(BIN) ./cmd/zebrapad/*.go

version: $(BIN)
	go version -m ./$(BIN)

run: $(BIN)
	./$(BIN)
