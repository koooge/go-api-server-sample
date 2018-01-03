GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

build:
	$(GO_BUILD_ENV) go build -v -o api-server ./...

run:
	go run src/server.go

clean:
	@rm -v api-server

.PHONY: build run clean
