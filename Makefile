all: test

test:
	go test ./... -v

cover:
	go test ./... -cover
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

deps:
	go mod tidy
