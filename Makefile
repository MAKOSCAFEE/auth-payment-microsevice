.PHONY: install test build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get ./...

test: install
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o payment_build ./payment

serve: build
	./payment_build

clean:
	rm ./payment_build
