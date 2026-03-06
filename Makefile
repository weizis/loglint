test:
	go test ./...

run:
	go run ./cmd/loginit ./...

run-example:
	go run ./cmd/loginit ./example

build:
	go build -o loglint ./cmd/loginit

install:
	go install ./cmd/loginit

tidy:
	go mod tidy