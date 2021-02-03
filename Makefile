BIN := json-stringify

ifdef update
	u=-u
endif

export GO111MODULE=on

.PHONY: all
all: bin/$(BIN)

bin/$(BIN): cli/*.go cmd/$(BIN)/*.go clean
	go build -ldflags="-s -w" -o bin/$(BIN) cmd/$(BIN)/main.go

.PHONY: deps
deps:
	go get ${u} -d ./...
	go mod tidy

.PHONY: test
test:
	go test -race ./...

.PHONY: clean
clean:
	rm -f bin/$(BIN)
