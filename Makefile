.PHONY: install test build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get ./...

test: install
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o ./payment/payment_build ./payment

docker_build: build
	docker build -t payment:v1 ./payment
	# docker build -t authentication:v1 ./auth
	# docker build -t summary:v1 ./summary
pack:
	GOOS=linux make build
	docker build -t gcr.io/barnie/payment:$(TAG) ./payment

serve: build
	./payment/payment_build

clean:
	rm ./payment/payment_build

upload:
	docker push gcr.io/barnie/payment:$(TAG)

k8s: docker_build pack upload
	kubectl apply -f ./k8s

