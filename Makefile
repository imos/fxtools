all:
	make format
	make build

format:
	go fmt ./...

build:
	go build cli/main.go

test:
	go test ./...
