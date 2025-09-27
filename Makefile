BIN=zebrapad

.PHONY: all clean test build version run
all: clean test build version run

clean:
	rm -rf ./bin/

go.mod:
	go mod init github.com/lvdh/zebrapad

test: go.mod
	go test -v ./...

$(BIN): build
build: go.mod cmd/zebrapad/*.go internal/* clean
	go build -o ./bin/$(BIN) ./cmd/zebrapad/*.go
	GOOS=windows GOARCH=amd64 go build -o bin/$(BIN).exe ./cmd/zebrapad/*.go
	du -sh ./bin/*

version: ./bin/
	go version -m ./bin/*

run: go.mod cmd/zebrapad/*.go internal/*
	@(sleep 1 && xdg-open "http://localhost:8080/" 2>/dev/null) &
	@(sleep 1 && echo ">> http://localhost:8080/") &
	go run ./...
