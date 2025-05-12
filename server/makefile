BINARY_NAME := yata-server

.DEFAULT_GOAL := dev


# run the server locally
dev:
	go run main.go

# run the built binary
preview: build
	dist/bin/$(BINARY_NAME)

# build the server into a binary
build:
	go build -o dist/bin/$(BINARY_NAME) .

# remove any build artifacts
clean:
	rm -rf ./dist

# format the code
fmt:
	go fmt .

# run the tests
test:
	go test ./...