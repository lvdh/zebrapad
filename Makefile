BIN=bin/zebrapad

.PHONY: all setup build test run clean

all: test $(BIN)

setup:
	cd src && go get github.com/gin-gonic/gin

build: $(BIN)
$(BIN): src/*.go src/web/*
	mkdir -p bin
	cd src && go build -o ../$(BIN) .

test:
	cd src && go test .

run: $(BIN)
	./bin/$(BIN)

clean:
	rm -rf bin/
