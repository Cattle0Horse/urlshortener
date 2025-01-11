FROM golang:1.23-alpine AS builder

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .

RUN go mod download -x
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .
COPY config.yaml .

ENV GIN_MODE=release

EXPOSE 8080/tcp

CMD ["./main"]
