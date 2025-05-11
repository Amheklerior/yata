BINARY_NAME := yata-server

.DEFAULT_GOAL := dev


dev:
	go run main.go

preview: build
	dist/bin/$(BINARY_NAME)

build:
	go build -o dist/bin/$(BINARY_NAME) .

clean:
	rm -rf ./dist

fmt:
	go fmt .

test:
	go test ./...