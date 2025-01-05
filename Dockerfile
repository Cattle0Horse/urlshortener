FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o shorturl ./cmd/server

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/shorturl .
COPY config.yaml .

EXPOSE 8080

ENTRYPOINT ["./shorturl"]
