mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/miniboard && \
	go mod tidy

run:
	cd cmd/miniboard/ && \
	go run .

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

check: fmt lint

go-build:
	cd cmd/miniboard/ && \
	CGO_ENABLED=0 go build -o ../../tmp/miniboard .