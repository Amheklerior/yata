BINARY_NAME := yata-server

.DEFAULT_GOAL := dev


dev:
	air

preview: build
	dist/bin/$(BINARY_NAME)

build:
	go build -o dist/bin/$(BINARY_NAME) .

clean:
	rm -rf ./dist

fmt:
	go fmt .