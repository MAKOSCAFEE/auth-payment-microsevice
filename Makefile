.PHONY: install test build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get ./...

test: install
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o payment_build ./payment

docker_build:
	docker build -t payment:v1 ./payment
	# docker build -t authentication:v1 ./auth
	# docker build -t summary:v1 ./summary

serve: build
	./payment_build

clean:
	rm ./payment_build

k8s: docker_build
	kubectl apply -f ./k8s