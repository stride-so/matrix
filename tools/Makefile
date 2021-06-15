test:
	go test -v ./... -short

.PHONY: test

clean:
	find ./tmp -type f -exec rm {} \+

.PHONY: clean

build:
	go build -o ./stride cmd/stride/main.go

