# uniapp for URLify

## 安装依赖

```shell
npm install --legacy-peer-dep
```

## 开发

```shell
npm run dev:h5
```

## 打包

```shell
npm run build:h5
```

## 打包后运行

如果需要打包成 `h5` 运行，可以使用 `http-server` 运行

```shell
# 安装 http-server
npm install -g http-server
# 运行在5173端口
http-server -p 5173 ./dist/build/h5
```

## 容器化编译

```shell
docker build -t uniapp-urlify .
```

```shell
docker run -it -p 5173:5173 uniapp-urlify
```

## Todo

1. 重定向逻辑处理，使用nginx部署后，点击登录并不会重定向至首页
2. 创建短链接后，自动刷新短链接列表
3. 短链接列表页显示到期时间
