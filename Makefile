run:
	go run ./cmd

build:
	go build -o bin/app ./cmd

lint:
	golangci-lint run

format:
	gofmt -w . && goimports -w .

test:
	go test ./...

