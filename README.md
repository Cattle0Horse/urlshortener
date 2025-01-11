# URLify

## TODO

- [x] 过期处理（redis中可以存储过期时间，或者redis中设置到期时间少于数据库过期时间） - 缓存数据库一致性
- [x] 添加布隆过滤器减少缓存击穿情况的发生
- [x] 容器化部署
- [ ] 链接访问统计
- [ ] 链接访问数据导出
- [ ] 批量生成短链接功能（如提交文件）
- [ ] 优化连接关闭（返回 clean 接口，类似于 `contextWithCancel`，可参考 [Apache-Answer](https://github.com/apache/incubator-answer) ）
- [ ] 可观测性监控
- [ ] 连接重试（如数据库），而非 panic（可参考 [Apache-Answer](https://github.com/apache/incubator-answer) ）
- [ ] DbProxy 数据库集群
- [ ] 读写分离（参考 [beihai0xff/turl](https://github.com/beihai0xff/turl) ）
- [ ] token 缓存
- [ ] token 有效期
- [ ] log 打印（参考 `internal/module/user/register.go`）
- [ ] 数据库存储可以考虑换成 base62编码 前的数字，这保证了有序性，数据库查询效率更优秀（不过这导致旧键无法复用，本来也不使用）
- [ ] 服务限流
- [ ] 健康检查
- [ ] 集成测试与单元测试，可以对于有依赖的内容可以借助 [testcontainers](https://github.com/testcontainers/testcontainers-go)

待考虑项：

- [ ] 相同 url 可以幂等
- [ ] 过期短链的处理（如轮询扫描全表，删除过期的）

## 目录设计说明

```text
├── cmd
│   ├── gen
│   │   └── gen.go # 依赖 `internal/model` 用于生成 `internal/query` 代码
│   └── server
│       └── server.go
├── config # golang配置包
├── deploy # 部署任务相关
|-- main.go #程序入口
├── internal
│   ├── global # 为内部提供全局变量或函数
│   │   ├── database # 数据相关
│   │   |   ├── mysql
│   │   |   └── redis
│   │   ├── logger # 日志相关
│   │   ├── query # gorm.io/gen 生成的数据库相关操作
│   │   └── middleware # 中间件
│   ├── module # 模块（或controller），如短链模块
│   └── model # 数据库模型
├── pkg # 公共包，最多依赖 config
│   └── tools # 直接函数，如异常处理，判断等（不需要init的工具，防止初始化影响其他包）
```

## 使用

```shell
# todo: 主容器等待依赖完全启动
docker compose -f ./deploy/docker-compose.yaml -p urlshortener-net up
```

## 开发

### 前端

```shell
cd ./frontend
pnpm install
pnpm dev
```

### 后端

可查看 `scripts/scripts.nu`

```shell
# 启动后端服务器
go run main.go
```
