FROM golang:1.23-alpine AS builder

RUN go env -w GOPROXY='https://goproxy.io,https://goproxy.cn,direct'

WORKDIR /app

COPY go.mod go.sum ./

# 清理go模块缓存
RUN go clean -modcache

# 下载依赖
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest AS production

# 安装ca-certificates管理CA证书，tzdata 管理时区信息
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app
COPY --from=builder /app/main .
COPY config.yaml .

ENV GIN_MODE=release

EXPOSE 8080

CMD ["./main"]
