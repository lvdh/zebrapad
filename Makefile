BIN=bin/zebrapad

.PHONY: all test run clean

all: test $(BIN)

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
