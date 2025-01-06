FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o shorturl.exe ./cmd/server

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/shorturl.exe .
COPY config.yaml .

EXPOSE 8080/tcp

ENTRYPOINT ["./shorturl"]
