############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
COPY . $GOPATH/src/auth-payment-microservice/auth
WORKDIR $GOPATH/src/auth-payment-microservice/auth
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/authentication
############################
# STEP 2 build a small image
############################

FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 5050

WORKDIR /root
# Copy our static executable.
COPY --from=builder /go/bin/authentication .


# Run the payment binary.
ENTRYPOINT ["./authentication"]