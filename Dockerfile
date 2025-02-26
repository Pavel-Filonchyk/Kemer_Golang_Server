    FROM golang:1.20 AS builder

    WORKDIR /app

    COPY go.mod ./
    COPY go.sum ./
    RUN go mod download

    COPY . .

    RUN go build -o myapp

    FROM alpine:latest

    WORKDIR /root/

    COPY --from=builder /app/myapp .

    CMD ["./myapp"]
