# syntax=docker/dockerfile:1

## build
FROM golang:1.18-alpine AS builder

WORKDIR /server
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o app .

## Deploy
FROM alpine:latest
WORKDIR /app
COPY --from=builder /server/app .
COPY init.sh ./init.sh
RUN chmod +x ./init.sh
ENTRYPOINT ["sh", "init.sh"]