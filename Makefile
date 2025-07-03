BIN=zebrapad

.PHONY: all clean test build run

all: clean test build run

go.mod:
	go mod init github.com/lvdh/zebrapad

clean:
	rm -rf ./$(BIN)

test: go.mod
	go test -v ./...

$(BIN): build
build: go.mod cmd/zebrapad/*.go internal/*
	go build -o ./$(BIN) ./cmd/zebrapad/*.go

run: go.mod $(BIN)
	./$(BIN)
