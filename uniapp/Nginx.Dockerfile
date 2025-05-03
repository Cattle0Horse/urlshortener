# 使用 Node.js 的官方镜像作为基础镜像
FROM node:22.11.0-alpine AS builder

# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk add --no-cache libc6-compat
# 设置工作目录
WORKDIR /app

# 复制 package.json 和 package-lock.json
COPY package*.json ./

# 安装项目依赖
RUN npm install --legacy-peer-deps

# 复制项目文件到工作目录
COPY . .

# 构建项目
RUN npm run build:h5

# Nginx 部署
FROM nginx:latest

COPY --from=builder /app/dist/build/h5 /app/uniapp

# 暴露端口
EXPOSE 80
