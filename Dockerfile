FROM golang:1.23.5-alpine AS builder
RUN apk add --no-cache git build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w" -o projectcinema-main && strip projectcinema-main

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/projectcinema-main .

EXPOSE 8086
CMD ["./projectcinema-main"]
