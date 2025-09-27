BIN=zebrapad

.PHONY: all clean test build version run
all: clean test build version run

clean:
	rm -rf ./$(BIN)

go.mod:
	go mod init github.com/lvdh/zebrapad

test: go.mod
	go test -v ./...

$(BIN): build
build: go.mod cmd/zebrapad/*.go internal/* clean
	go build -o ./$(BIN) ./cmd/zebrapad/*.go
	du -sh ./$(BIN)

version: $(BIN)
	go version -m ./$(BIN)

run: go.mod cmd/zebrapad/*.go internal/*
	@(sleep 1 && xdg-open "http://localhost:8080/" 2>/dev/null) &
	@(sleep 1 && echo ">> http://localhost:8080/") &
	go run ./...
