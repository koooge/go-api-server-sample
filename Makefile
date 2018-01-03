build:
	go build -o api-server

run:
	go run server.go

clean:
	@rm -v api-server

.PHONY: build run clean
