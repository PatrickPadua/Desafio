ARG GO_VERSION=1.12
ARG ALPINE_VERSION=3.10.1

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o worker

FROM alpine:${ALPINE_VERSION}

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/worker .

EXPOSE 8080

ENTRYPOINT ["./worker"]