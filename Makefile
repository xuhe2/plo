.PHONY: build run test clean

build:
	mkdir -p bin
	go build -o bin/plo main.go

run:
	go run main.go -input example.drawio -output output.md

test:
	go test ./...

clean:
	rm -rf bin output.md
