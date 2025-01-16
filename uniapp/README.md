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
