FROM golang:1.18-alpine3.16 AS builder

ENV CGO_ENABLED=1 GOOS=linux
RUN apk add --no-cache \
    gcc \
    musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

FROM alpine:3.16

COPY /cmd/ .
COPY /pkg/ .
COPY --from=builder /app/ .

EXPOSE 8081
CMD ["./main"]