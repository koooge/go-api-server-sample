build:
	go build -v -o api-server ./...

run:
	go run src/server.go

clean:
	@rm -v api-server

.PHONY: build run clean
