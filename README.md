## TODO

- [x] 添加布隆过滤器减少缓存击穿情况的发生
- [ ] Docker 部署
- [ ] 优化连接关闭（返回 clean 接口，类似于 `contextWithCancel`，可参考 [Apache-Answer](https://github.com/apache/incubator-answer) ）
- [ ] 可观测性监控
- [ ] 连接重试（如数据库），而非 panic（可参考 [Apache-Answer](https://github.com/apache/incubator-answer) ）
- [ ] DbProxy 数据库集群
- [ ] 读写分离（参考 [beihai0xff/turl](https://github.com/beihai0xff/turl) ）
- [ ] token 缓存
- [ ] 过期处理（redis中可以存储过期时间，或者redis中设置到期时间少于数据库过期时间） - 缓存数据库一致性
- [ ] log 打印（参考 `internal/module/user/register.go`）
- [ ] 数据库存储可以考虑换成 base62编码 前的数字，这保证了有序性，数据库查询更优秀（不过这导致旧键无法复用，本来也不使用）

考虑项：

- [ ] 相同 url 可以幂等，思考是否需要幂等？
- [ ] 过期短链的处理（如轮询扫描全表，删除过期的） - 是否需要处理？

## 目录设计说明

`pkg/tools`:
`config`: 配置文件，初始化最高优先级
`cmd/gen`: `gorm` 代码生成
`cmd/server`: 服务启动
`internal/global`: 提供全局可访问的，大多依赖 `config` 和 `pkg` 包
`internal/`

```text
├── cmd
│   ├── gen
│   │   └── gen.go # 依赖 `internal/model`
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
