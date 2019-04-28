all: test build run

test:
	go test ./...
build:
	go build main.go 
run:
	./main 10